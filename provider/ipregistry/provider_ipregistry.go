package ipregistry

import (
	"errors"

	"github.com/evalphobia/ipregistry-go/config"
	"github.com/evalphobia/ipregistry-go/ipregistry"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPRegistryProvider struct {
	client *ipregistry.IPRegistry
}

func (p *IPRegistryProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPRegistryProvider")
	}

	svc, err := ipregistry.New(config.Config{
		APIKey: c.IPRegistryAPIKey,
		Debug:  c.Debug,
	})
	if err != nil {
		return err
	}

	p.client = svc
	return nil
}

func (p IPRegistryProvider) String() string {
	return "ipregistry"
}

func (p IPRegistryProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.SingleIP(ipaddr)
	if err != nil {
		return emptyResult, err
	}
	if resp.HasError() {
		return emptyResult, errors.New(resp.ErrData.Message)
	}

	loc := resp.Location
	sec := resp.Security
	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IP,
		Hostname:       resp.Hostname,
		Organization:   resp.Connection.Organization,
		ASNumber:       resp.Connection.ASN,
		IsHosting:      sec.IsCloudProvider,
		IsProxy:        sec.IsProxy,
		IsTor:          sec.IsTor || sec.IsTorExit,
		IsBogon:        sec.IsBogon,
		HasOtherThreat: sec.IsThreat,
		Country:        loc.Country.Code,
		City:           loc.City,
		Region:         loc.Region.Name,
		Latitude:       loc.Latitude,
		Longitude:      loc.Longitude,
	}, nil
}

func (p IPRegistryProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.SingleIP(ipaddr)
}
