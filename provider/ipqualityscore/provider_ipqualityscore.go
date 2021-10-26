package ipqualityscore

import (
	"errors"

	"github.com/evalphobia/ipqualityscore-go/config"
	"github.com/evalphobia/ipqualityscore-go/ipqs"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPQualityScoreProvider struct {
	client *ipqs.IPQualityScore
}

func (p *IPQualityScoreProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPQualityScoreProvider")
	}

	svc, err := ipqs.New(config.Config{
		APIKey: c.IPQualityScoreAPIKey,
		Debug:  c.Debug,
	})
	if err != nil {
		return err
	}

	p.client = svc
	return nil
}

func (p IPQualityScoreProvider) String() string {
	return "ipqualityscore"
}

func (p IPQualityScoreProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.IPReputation(ipaddr)
	if err != nil {
		return emptyResult, err
	}
	if resp.HasError() {
		return emptyResult, errors.New(resp.ErrData.Message)
	}

	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             ipaddr,
		Hostname:       resp.Host,
		Organization:   resp.Organization,
		RiskScore:      resp.FraudScore / float64(100),
		ASNumber:       resp.ASN,
		IsProxy:        resp.Proxy,
		IsAnonymousVPN: resp.VPN || resp.ActiveVPN,
		IsTor:          resp.Tor || resp.ActiveTor,
		IsBot:          resp.BotStatus,
		HasOtherThreat: resp.RecentAbuse,
		Country:        resp.CountryCode,
		City:           resp.City,
		Region:         resp.Region,
		Latitude:       resp.Latitude,
		Longitude:      resp.Longitude,
	}, nil
}

func (p IPQualityScoreProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.IPReputation(ipaddr)
}
