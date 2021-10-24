package ipinfoio

import (
	"errors"
	"net"
	"strconv"
	"strings"

	"github.com/ipinfo/go/v2/ipinfo"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPinfoioProvider struct {
	client *ipinfo.Client
}

func (p *IPinfoioProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPinfoioProvider")
	}
	token := c.GetIPinfoioToken()
	if token == "" {
		return errors.New("empty token for IPinfoioProvider")
	}

	p.client = ipinfo.NewClient(nil, nil, token)
	return nil
}

func (p IPinfoioProvider) String() string {
	return "ipinfo.io"
}

func (p IPinfoioProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.GetIPInfo(net.ParseIP(ipaddr))
	if err != nil {
		return emptyResult, err
	}

	var lat, longi float64
	loc := strings.Split(resp.Location, ",")
	if len(loc) == 2 {
		lat, _ = strconv.ParseFloat(loc[0], 64)
		longi, _ = strconv.ParseFloat(loc[1], 64)
	}

	var privacy ipinfo.CorePrivacy
	if resp.Privacy != nil {
		privacy = *resp.Privacy
	}

	var as ipinfo.CoreASN
	var asn int64
	if resp.ASN != nil {
		as = *resp.ASN
		asn, _ = strconv.ParseInt(strings.TrimPrefix(as.ASN, "AS"), 10, 64)
	}

	org := resp.Org
	if resp.ASN == nil && strings.HasPrefix(org, "AS") {
		parts := strings.Split(org, " ")
		if len(parts) > 1 {
			num, err := strconv.ParseInt(strings.TrimPrefix(parts[0], "AS"), 10, 64)
			if err == nil {
				asn = num
				org = strings.Join(parts[1:], " ")
			}
		}
	}

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IP.String(),
		Hostname:       resp.Hostname,
		ISP:            as.Name,
		Organization:   org,
		ASNumber:       asn,
		IsAnonymousVPN: privacy.VPN,
		IsHosting:      privacy.Hosting,
		IsProxy:        privacy.Proxy,
		IsTor:          privacy.Tor,
		IsBot:          false,
		IsBogon:        resp.Bogon,
		Region:         resp.Region,
		Country:        resp.Country,
		City:           resp.City,
		Latitude:       lat,
		Longitude:      longi,
	}, nil
}

func (p IPinfoioProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.GetIPInfo(net.ParseIP(ipaddr))
}
