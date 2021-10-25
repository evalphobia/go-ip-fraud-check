package main

import (
	"strconv"
	"strings"
	"sync"

	"github.com/mkideal/cli"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/log"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

var outputHeader = []string{
	"service",
	"ip_address",
	"hostname",
	"risk_score",
	"isp",
	"organization",
	"asn",
	"country",
	"city",
	"region",
	"latitude",
	"longitude",
	"is_anonymous",
	"is_anonymous_vpn",
	"is_hosting",
	"is_proxy",
	"is_tor",
	"is_bot",
	"is_bogon",
	"has_other_threat",
	// "as_routes", append by code
}

// parameters of 'list' command.
type listT struct {
	cli.Helper
	Provider string `cli:"*p,provider" usage:"set types of api provider (space separated) --provider='ipdata ipinfo minfraud'"`
	InputCSV string `cli:"*i,input" usage:"input csv/tsv file path --input='./input.csv'"`
	Output   string `cli:"*o,output" usage:"output tsv file path --output='./output.tsv'"`
	UseRoute bool   `cli:"route" usage:"set if you need route data from IRR (this might be slow) --route"`
	Debug    bool   `cli:"debug" usage:"set if you use HTTP debug feature --debug"`
}

func (a *listT) Validate(ctx *cli.Context) error {
	return validateProviderString(a.Provider)
}

var listC = &cli.Command{
	Name: "list",
	Desc: "Exec api call of ip address fraud check providers from csv list file",
	Argv: func() interface{} { return new(listT) },
	Fn:   execList,
}

func execList(ctx *cli.Context) error {
	argv := ctx.Argv().(*listT)

	r := newListRunner(*argv)
	return r.Run()
}

type ListRunner struct {
	// parameters
	Provider string
	InputCSV string
	Output   string
	UseRoute bool
	Debug    bool
}

func newListRunner(p listT) ListRunner {
	return ListRunner{
		Provider: p.Provider,
		InputCSV: p.InputCSV,
		Output:   p.Output,
		UseRoute: p.UseRoute,
		Debug:    p.Debug,
	}
}

func (r *ListRunner) Run() error {
	f, err := NewCSVHandler(r.InputCSV)
	if err != nil {
		return err
	}

	w, err := NewFileHandler(r.Output)
	if err != nil {
		return err
	}

	lines, err := f.ReadAll()
	if err != nil {
		return err
	}

	maxReqNum := 3
	maxReq := make(chan struct{}, maxReqNum)

	providerList, err := getProvidersFromString(r.Provider)
	if err != nil {
		panic(err)
	}

	logger := &log.StdLogger{}
	svc, err := ipfraudcheck.New(ipfraudcheck.Config{
		UseRoute: r.UseRoute,
		Debug:    r.Debug,
		Logger:   logger,
	}, providerList)
	if err != nil {
		panic(err)
	}
	if r.UseRoute {
		outputHeader = append(outputHeader, "as_routes")
	}

	providerSize := len(providerList)
	result := make([]string, len(lines)*providerSize)
	var wg sync.WaitGroup
	for i, line := range lines {
		i = i * providerSize
		wg.Add(1)
		go func(i int, line map[string]string) {
			maxReq <- struct{}{}
			defer func() {
				<-maxReq
				wg.Done()
			}()

			logger.Infof("exec #: [%d]\n", i/providerSize)
			rows, err := r.execAPI(svc, line)
			if err != nil {
				logger.Errorf("#: [%d]; err=[%v]\n", i, err)
				return
			}
			for j, row := range rows {
				result[i+j] = strings.Join(row, "\t")
			}
		}(i, line)
	}
	wg.Wait()

	result = append([]string{strings.Join(outputHeader, "\t")}, result...)
	return w.WriteAll(result)
}

func (r *ListRunner) execAPI(svc *ipfraudcheck.Client, param map[string]string) ([][]string, error) {
	resp, err := svc.CheckIP(param["ip_address"])
	if err != nil {
		return nil, err
	}

	rows := make([][]string, len(resp.List))
	for i, r := range resp.List {
		row := make([]string, 0, len(outputHeader))
		for _, v := range outputHeader {
			if v == "as_routes" {
				row = append(row, strings.Join(resp.ASPrefix, " "))
				continue
			}
			row = append(row, getValue(param, r, v))
		}
		rows[i] = row
	}
	return rows, nil
}

func getValue(param map[string]string, resp provider.FraudCheckResponse, name string) string {
	switch name {
	case "service":
		return resp.ServiceName
	case "ip_address":
		return resp.IP
	case "hostname":
		return resp.Hostname
	case "risk_score":
		return strconv.FormatFloat(resp.RiskScore, 'f', 5, 64)
	case "isp":
		return resp.ISP
	case "organization":
		return resp.Organization
	case "asn":
		return strconv.FormatInt(resp.ASNumber, 10)
	case "country":
		return resp.Country
	case "city":
		return resp.City
	case "region":
		return resp.Region
	case "latitude":
		return strconv.FormatFloat(resp.Latitude, 'f', 5, 64)
	case "longitude":
		return strconv.FormatFloat(resp.Longitude, 'f', 5, 64)
	case "is_anonymous":
		return strconv.FormatBool(resp.IsAnonymous)
	case "is_anonymous_vpn":
		return strconv.FormatBool(resp.IsAnonymousVPN)
	case "is_hosting":
		return strconv.FormatBool(resp.IsHosting)
	case "is_proxy":
		return strconv.FormatBool(resp.IsProxy)
	case "is_tor":
		return strconv.FormatBool(resp.IsTor)
	case "is_bot":
		return strconv.FormatBool(resp.IsBot)
	case "is_bogon":
		return strconv.FormatBool(resp.IsBogon)
	case "has_other_threat":
		return strconv.FormatBool(resp.HasOtherThreat)
	}
	return ""
}
