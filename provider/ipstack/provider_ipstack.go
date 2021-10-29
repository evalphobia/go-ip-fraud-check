package ipstack

import (
	"errors"
	"fmt"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPStackProvider struct {
	client Client
}

func (p *IPStackProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPStackProvider")
	}

	apiKey := c.GetIPStackAPIKey()
	if apiKey == "" {
		return errors.New("apikey for ipstack is empty. you must set directly or use 'FRAUD_CHECK_IPSTACK_APIKEY' envvar")
	}
	cli := NewClient(apiKey)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p IPStackProvider) String() string {
	return "ipstack"
}

func (p IPStackProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.LookUpWithSecurity(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	conn := resp.Connection
	sec := resp.Security
	score := float64(0)
	switch sec.ThreatLevel {
	case "middle":
		score = 0.5
	case "high":
		score = 0.75
	}
	comment := ""
	switch {
	case len(sec.ThreatTypes) != 0 && sec.CrawlerType != "":
		comment = fmt.Sprintf("threats:%+v, crawler_type:%+v", sec.ThreatTypes, sec.CrawlerType)
	case len(sec.ThreatTypes) != 0:
		comment = fmt.Sprintf("threats:%+v", sec.ThreatTypes)
	case sec.CrawlerType != "":
		comment = fmt.Sprintf("crawler_type:%s", sec.CrawlerType)
	}

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IP,
		Hostname:       resp.Hostname,
		ISP:            conn.ISP,
		ASNumber:       conn.ASN,
		Country:        resp.CountryCode,
		City:           resp.City,
		Region:         resp.RegionName,
		Latitude:       resp.Latitude,
		Longitude:      resp.Longitude,
		RiskScore:      score,
		IsProxy:        sec.IsProxy && sec.ProxyType != "vpn",
		IsVPN:          sec.ProxyType == "vpn",
		IsTor:          sec.IsTor,
		HasOtherThreat: sec.IsCrawler,
		ThreatComment:  comment,
	}, nil
}

func (p IPStackProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.LookUpWithSecurity(ipaddr)
}
