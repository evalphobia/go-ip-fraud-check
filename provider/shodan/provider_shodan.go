package shodan

import (
	"errors"
	"fmt"
	"strconv"
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type ShodanProvider struct {
	client Client
}

func (p *ShodanProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for ShodanProvider")
	}

	apiKey := c.GetShodanAPIKey()
	if apiKey == "" {
		return errors.New("apikey for Shodan is empty. you must set directly or use 'FRAUD_CHECK_SHODAN_APIKEY' envvar")
	}
	cli := NewClient(apiKey)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p ShodanProvider) String() string {
	return "shodan"
}

func (p ShodanProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.GetHost(ipaddr)
	if err != nil {
		return emptyResult, err
	}
	asn, _ := strconv.ParseInt(strings.TrimPrefix(resp.ASN, "AS"), 10, 64)

	tags := make(map[string]struct{})
	for _, v := range resp.Tags {
		tags[v] = struct{}{}
	}
	ports := make(map[int]struct{})
	for _, v := range resp.Ports {
		ports[v] = struct{}{}
	}

	data := resp.Data
	_, isTor := tags["tor"]
	_, isVPN := tags["vpn"]
	if !isVPN {
		for _, v := range data {
			if v.ISAKMP != nil {
				isVPN = true
				break
			}
		}
	}
	_, isHosting := ports[22]
	if !isHosting {
		_, isHosting = tags["cloud"]
	}
	if !isHosting {
		for _, v := range data {
			if v.Cloud != nil || v.SSH != nil {
				isHosting = true
				break
			}
		}
	}

	comment := ""
	switch {
	case len(tags) != 0 && len(ports) != 0:
		comment = fmt.Sprintf("tags:%+v, ports:%+v", resp.Tags, resp.Ports)
	case len(tags) != 0:
		comment = fmt.Sprintf("tags:%+v", resp.Tags)
	case len(ports) != 0:
		comment = fmt.Sprintf("ports:%+v", resp.Ports)
	}

	return provider.FraudCheckResponse{
		ServiceName:   p.String(),
		IP:            ipaddr,
		ISP:           resp.ISP,
		Organization:  resp.Org,
		ASNumber:      asn,
		IsVPN:         isVPN,
		IsHosting:     isHosting,
		IsTor:         isTor,
		ThreatComment: comment,
		Country:       resp.CountryCode,
		City:          resp.City,
		Latitude:      resp.Latitude,
		Longitude:     resp.Latitude,
	}, nil
}

func (p ShodanProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.GetHost(ipaddr)
}
