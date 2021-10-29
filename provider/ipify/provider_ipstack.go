package ipify

import (
	"errors"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPifyProvider struct {
	client Client
}

func (p *IPifyProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPifyProvider")
	}

	apiKey := c.GetIPifyAPIKey()
	if apiKey == "" {
		return errors.New("apikey for ipify is empty. you must set directly or use 'FRAUD_CHECK_IPIFY_APIKEY' envvar")
	}
	cli := NewClient(apiKey)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p IPifyProvider) String() string {
	return "ipify"
}

func (p IPifyProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.LookUpWithSecurity(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	as := resp.AS
	loc := resp.Location
	sec := resp.Proxy
	return provider.FraudCheckResponse{
		ServiceName: p.String(),
		IP:          resp.IP,
		ISP:         resp.ISP,
		ASNumber:    as.ASN,
		Country:     loc.Country,
		City:        loc.City,
		Region:      loc.Region,
		Latitude:    loc.Latitude,
		Longitude:   loc.Longitude,
		IsProxy:     sec.Proxy,
		IsVPN:       sec.VPN,
		IsTor:       sec.Tor,
	}, nil
}

func (p IPifyProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.LookUpWithSecurity(ipaddr)
}
