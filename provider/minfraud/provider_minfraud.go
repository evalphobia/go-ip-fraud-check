package minfraud

import (
	"errors"

	"github.com/evalphobia/minfraud-api-go/config"
	"github.com/evalphobia/minfraud-api-go/minfraud"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type MinFraudProvider struct {
	client *minfraud.MinFraud
}

func (p *MinFraudProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for MinFraudProvider")
	}

	svc, err := minfraud.New(config.Config{
		AccountID:  c.MinFraudAccountID,
		LicenseKey: c.MinFraudLicenseKey,
		Debug:      c.Debug,
	})
	if err != nil {
		return err
	}

	p.client = svc
	return nil
}

func (p MinFraudProvider) String() string {
	return "minFraud"
}

func (p MinFraudProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.InsightsByIP(ipaddr)
	if err != nil {
		return emptyResult, err
	}
	if resp.HasError() {
		return emptyResult, errors.New(resp.ErrData.Error)
	}

	ip := resp.IPAddress
	traits := ip.Traits
	return provider.FraudCheckResponse{
		ServiceName:  p.String(),
		IP:           traits.IPAddress,
		Hostname:     traits.Domain,
		ISP:          traits.ISP,
		Organization: traits.Organization,
		ASNumber:     traits.AutonomousSystemNumber,
		RiskScore:    ip.Risk,
		IsVPN:        traits.IsAnonymousVPN,
		IsHosting:    traits.IsHostingProvider,
		IsProxy:      traits.IsPublicProxy || traits.IsResidentialProxy,
		IsTor:        traits.IsTorExitNode,
		IsBot:        false,
		Country:      ip.Country.ISOCode,
		City:         ip.City.Names.EN,
		Latitude:     ip.Location.Latitude,
		Longitude:    ip.Location.Longitude,
	}, nil
}

func (p MinFraudProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.InsightsByIP(ipaddr)
}
