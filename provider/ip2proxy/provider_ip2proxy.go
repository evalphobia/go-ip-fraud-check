package ip2proxy

import (
	"errors"
	"strconv"

	"github.com/ip2location/ip2proxy-go/v3"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IP2ProxyProvider struct {
	client *ip2proxy.WS
}

func (p *IP2ProxyProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IP2ProxyProvider")
	}

	apiKey := c.GetIP2ProxyAPIKey()
	if apiKey == "" {
		return errors.New("apikey for ip2proxy is empty. you must set directly or use 'FRAUD_CHECK_IP2PROXY_APIKEY' envvar")
	}
	svc, err := ip2proxy.OpenWS(apiKey, c.GetIP2ProxyAPIPackage(), true)
	if err != nil {
		return err
	}
	p.client = svc
	return nil
}

func (p IP2ProxyProvider) String() string {
	return "ip2proxy"
}

func (p IP2ProxyProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.LookUp(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	asn, _ := strconv.ParseInt(resp.ASN, 10, 64)

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             ipaddr,
		Hostname:       resp.Domain,
		ISP:            resp.ISP,
		ASNumber:       asn,
		Organization:   resp.Provider,
		IsProxy:        resp.IsProxy == "YES",
		IsVPN:          resp.ProxyType == "VPN",
		IsHosting:      resp.ProxyType == "DCH",
		IsTor:          resp.ProxyType == "TOR",
		IsBot:          resp.ProxyType == "SES",
		HasOtherThreat: resp.Threat != "",
		ThreatComment:  resp.Threat,
		Region:         resp.RegionName,
		Country:        resp.CountryCode,
		City:           resp.CityName,
	}, nil
}

func (p IP2ProxyProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.LookUp(ipaddr)
}
