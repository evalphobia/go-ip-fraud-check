package ipapicom

import (
	"errors"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPAPIComProvider struct {
	client Client
}

func (p *IPAPIComProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPAPIComProvider")
	}

	cli := NewClient(c.GetIPAPIComAPIKey())
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p IPAPIComProvider) String() string {
	return "ip-api"
}

func (p IPAPIComProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.LookUpWithSecurity(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	return provider.FraudCheckResponse{
		ServiceName:  p.String(),
		IP:           ipaddr,
		ISP:          resp.ISP,
		Organization: resp.Org,
		ASNumber:     resp.ASNumber(),
		Country:      resp.Country,
		City:         resp.City,
		Region:       resp.Region,
		Latitude:     resp.Latitude,
		Longitude:    resp.Longitude,
		IsProxy:      resp.Proxy,
		IsHosting:    resp.Hosting,
	}, nil
}

func (p IPAPIComProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.LookUpWithSecurity(ipaddr)
}
