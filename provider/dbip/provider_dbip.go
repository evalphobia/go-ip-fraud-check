package dbip

import (
	"errors"
	"fmt"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type DBIPProvider struct {
	client Client
}

func (p *DBIPProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for DBIPProvider")
	}

	apiKey := c.GetDBIPAPIKey()
	if apiKey == "" {
		return errors.New("apikey for dbip is empty. you must set directly or use 'FRAUD_CHECK_DBIP_APIKEY' envvar")
	}
	cli := NewClient(apiKey)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p DBIPProvider) String() string {
	return "dbip"
}

func (p DBIPProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.LookUp(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	comment := ""
	switch {
	case len(resp.ThreatDetails) != 0:
		comment = fmt.Sprintf("threats:%+v", resp.ThreatDetails)
	}

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IPAddress,
		ISP:            resp.ISP,
		Organization:   resp.Organization,
		ASNumber:       resp.ASNumber,
		Country:        resp.CountryCode,
		City:           resp.City,
		Latitude:       resp.Latitude,
		Longitude:      resp.Longitude,
		RiskScore:      resp.GetRiskScore(),
		IsProxy:        resp.IsHTTPProxy(),
		IsVPN:          resp.IsVPN(),
		IsHosting:      resp.IsHosting(),
		IsTor:          resp.IsTor(),
		IsBot:          resp.IsBot(),
		HasOtherThreat: resp.HasOtherThreat(),
		ThreatComment:  comment,
	}, nil
}

func (p DBIPProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.LookUp(ipaddr)
}
