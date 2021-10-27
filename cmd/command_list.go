package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"strings"
	"sync"
	"syscall"
	"time"

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
	"is_vpn",
	"is_hosting",
	"is_proxy",
	"is_tor",
	"is_bot",
	"is_bogon",
	"has_other_threat",
	"threat_comment",
	"error",
	// "as_routes", append by code
}

// parameters of 'list' command.
type listT struct {
	cli.Helper
	Provider string `cli:"*p,provider" usage:"set types of api provider (space separated) --provider='ipdata ipinfo minfraud'"`
	InputCSV string `cli:"*i,input" usage:"input csv/tsv file path --input='./input.csv'"`
	Output   string `cli:"*o,output" usage:"output tsv file path --output='./output.tsv'"`
	UseRoute bool   `cli:"route" usage:"set if you need route data from IRR (this might be slow) --route"`
	Interval string `cli:"interval" usage:"time interval after a API call to handle rate limit (ms=msec s=sec, m=min) --interval=1.5s"`
	Parallel int    `cli:"m,parallel" usage:"parallel number (multiple API calls) --parallel=2" dft:"2"`
	Verbose  bool   `cli:"v,verbose" usage:"set if you need detail logs --verbose"`
	Debug    bool   `cli:"debug" usage:"set if you use HTTP debug feature --debug"`
}

func (a *listT) Validate(ctx *cli.Context) error {
	if a.Interval != "" {
		if _, err := time.ParseDuration(a.Interval); err != nil {
			return fmt.Errorf("invalid 'interval' format: [%w]", err)
		}
	}

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
	Parallel int
	UseRoute bool
	Interval string
	Verbose  bool
	Debug    bool

	logger log.Logger
}

func newListRunner(p listT) ListRunner {
	return ListRunner{
		Provider: p.Provider,
		InputCSV: p.InputCSV,
		Output:   p.Output,
		Parallel: p.Parallel,
		UseRoute: p.UseRoute,
		Interval: p.Interval,
		Verbose:  p.Verbose,
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

	providerList, err := getProvidersFromString(r.Provider)
	if err != nil {
		panic(err)
	}

	// parse interval duration
	var interval time.Duration
	if r.Interval != "" {
		v, err := time.ParseDuration(r.Interval)
		if err != nil {
			return err
		}
		interval = v
	}

	logger := &log.StdLogger{}
	r.logger = logger
	svc, err := ipfraudcheck.New(ipfraudcheck.Config{
		UseRoute: r.UseRoute,
		Interval: interval,
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

	// handle kill signal
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-sig
		logger.Errorf("Stop signal detected!")
		logger.Errorf("Saving intermediate results...")
		result = append([]string{strings.Join(outputHeader, "\t")}, result...)
		w.WriteAll(result)
		os.Exit(2)
	}()

	maxReq := make(chan struct{}, r.Parallel)
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

			num := i / providerSize
			logger.Infof("exec #: [%d]\n", num)
			rows, err := r.execAPI(svc, line, num)
			if err != nil {
				logger.Errorf("#: [%d]; err=[%v]\n", i, err)
				return
			}
			for j, row := range rows {
				result[i+j] = strings.Join(row, "\t")
			}
			svc.WaitInterval()
		}(i, line)
	}
	wg.Wait()

	result = append([]string{strings.Join(outputHeader, "\t")}, result...)
	logger.Infof("Finished")
	return w.WriteAll(result)
}

func (r *ListRunner) execAPI(svc *ipfraudcheck.Client, param map[string]string, num int) ([][]string, error) {
	resp, err := svc.CheckIP(param["ip_address"])
	if err != nil {
		return nil, err
	}

	rows := make([][]string, len(resp.List))
	for i, res := range resp.List {
		row := make([]string, 0, len(outputHeader))
		for _, v := range outputHeader {
			if v == "as_routes" {
				row = append(row, strings.Join(resp.ASPrefix, " "))
				continue
			}
			row = append(row, getValue(param, res, v))
		}
		if r.Verbose {
			switch {
			case res.Err != "":
				r.logger.Errorf("#: [%d]; provider=[%s] ip=[%s]  err=[%v]\n", num, res.ServiceName, res.IP, res.Err)
			default:
				r.logger.Infof("#: [%d]; provider=[%s] ip=[%s]  row=[%v]\n", num, res.ServiceName, res.IP, row)
			}
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
	case "is_anonymous_vpn":
		return strconv.FormatBool(resp.IsVPN)
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
	case "threat_comment":
		return resp.ThreatComment
	case "error":
		return resp.Err
	}
	return ""
}
