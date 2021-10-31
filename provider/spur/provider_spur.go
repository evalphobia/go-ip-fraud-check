package spur

import (
	"errors"
	"fmt"
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type SpurProvider struct {
	client Client
}

func (p *SpurProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for SpurProvider")
	}

	token := c.GetSpurToken()
	if token == "" {
		return errors.New("token for spur is empty. you must set directly or use 'FRAUD_CHECK_SPUR_TOKEN' envvar")
	}
	cli := NewClient(token)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p SpurProvider) String() string {
	return "spur"
}

func (p SpurProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.DoContext(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	var comments []string
	hostingLoc := resp.GeoLite.Country
	geop := resp.GeoPrecision
	if geop.Exists {
		if hostingLoc != geop.Country {
			comments = append(comments, fmt.Sprintf("hosting_country=[%s] actual_country=[%s]", hostingLoc, geop.Country))
		}
	}
	if resp.Devices.Estimate >= 25 && !resp.IsMobile() {
		comments = append(comments, fmt.Sprintf("estimate_devices=[%d] infra=[%s]", resp.Devices.Estimate, resp.Infrastructure))
	}
	hasThreat := len(comments) != 0

	if resp.IsFileSharing() {
		comments = append(comments, "file_sharing")
	}
	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IP,
		Organization:   resp.AS.Organization,
		ASNumber:       resp.AS.Number,
		Country:        resp.GeoLite.Country,
		City:           resp.GeoLite.City,
		Latitude:       resp.GeoPrecision.Point.Latitude,
		Longitude:      resp.GeoPrecision.Point.Longitude,
		IsProxy:        resp.IsProxy(),
		IsVPN:          resp.IsVPN(),
		IsTor:          resp.IsTor(),
		IsHosting:      resp.IsHosting(),
		HasOtherThreat: hasThreat,
		ThreatComment:  strings.Join(comments, " | "),
	}, nil
}

func (p SpurProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.DoContext(ipaddr)
}
