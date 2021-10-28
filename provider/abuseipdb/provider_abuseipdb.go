package abuseipdb

import (
	"errors"
	"fmt"
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type AbuseIPDBProvider struct {
	client Client
}

func (p *AbuseIPDBProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for AbuseIPDBProvider")
	}

	apiKey := c.GetAbuseUPDBAPIKey()
	if apiKey == "" {
		return errors.New("apikey for AbuseIPDB is empty. you must set directly or use 'FRAUD_CHECK_ABUSEIPDB_APIKEY' envvar")
	}
	cli := NewClient(apiKey)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p AbuseIPDBProvider) String() string {
	return "AbuseIPDB"
}

func (p AbuseIPDBProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.Check(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	d := resp.Data
	var comments []string
	if d.NumDistinctUsers != 0 {
		comments = append(comments, fmt.Sprintf("distinct_users=[%d]", d.NumDistinctUsers))
	}
	if d.UsageType != "" {
		comments = append(comments, d.UsageType)
	}

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             ipaddr,
		Hostname:       d.Domain,
		ISP:            d.ISP,
		Country:        d.CountryCode,
		HasOtherThreat: d.TotalReports > 0,
		IsHosting:      d.IsHosting(),
		RiskScore:      float64(d.AbuseConfidenceScore) / 100,
		ThreatComment:  strings.Join(comments, " | "),
	}, nil
}

func (p AbuseIPDBProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.Check(ipaddr)
}
