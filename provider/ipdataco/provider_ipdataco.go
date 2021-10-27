package ipdataco

import (
	"errors"
	"strconv"
	"strings"

	ipdata "github.com/ipdata/go"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPdatacoProvider struct {
	client ipdata.Client
}

func (p *IPdatacoProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPdatacoProvider")
	}

	apiKey := c.GetIPdatacoAPIKey()
	if apiKey == "" {
		return errors.New("apikey for ipdata.co is empty. you must set directly or use 'FRAUD_CHECK_IPDATACO_APIKEY' envvar")
	}

	svc, err := ipdata.NewClient(apiKey)
	if err != nil {
		return err
	}
	p.client = svc
	return nil
}

func (p IPdatacoProvider) String() string {
	return "ipdata.co"
}

func (p IPdatacoProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.Lookup(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	var threat ipdata.Threat
	if resp.Threat != nil {
		threat = *resp.Threat
	}
	asn, _ := strconv.ParseInt(strings.TrimPrefix(resp.ASN.ASN, "AS"), 10, 64)

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IP,
		ISP:            resp.ASN.Name,
		Organization:   resp.Organization,
		ASNumber:       asn,
		IsProxy:        threat.IsProxy,
		IsTor:          threat.IsTOR,
		IsBogon:        threat.IsBogon,
		HasOtherThreat: threat.IsThreat,
		Region:         resp.Region,
		Country:        resp.CountryCode,
		City:           resp.City,
		Latitude:       resp.Latitude,
		Longitude:      resp.Longitude,
	}, nil
}

func (p IPdatacoProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.Lookup(ipaddr)
}
