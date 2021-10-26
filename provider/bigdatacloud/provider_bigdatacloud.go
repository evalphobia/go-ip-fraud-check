package bigdatacloud

import (
	"errors"

	"github.com/evalphobia/bigdatacloud-go/bigdatacloud"
	"github.com/evalphobia/bigdatacloud-go/config"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type BigDataCloudProvider struct {
	client *bigdatacloud.BigDataCloud
}

func (p *BigDataCloudProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for BigDataCloudProvider")
	}

	svc, err := bigdatacloud.New(config.Config{
		APIKey: c.BigDataCloudAPIKey,
		Debug:  c.Debug,
	})
	if err != nil {
		return err
	}

	p.client = svc
	return nil
}

func (p BigDataCloudProvider) String() string {
	return "bigdatacloud"
}

func (p BigDataCloudProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.IPGeolocationFull(ipaddr)
	if err != nil {
		return emptyResult, err
	}
	if resp.HasError() {
		return emptyResult, errors.New(resp.ErrData.Description)
	}

	var carrier bigdatacloud.Carrier
	if len(resp.Network.Carriers) != 0 {
		carrier = resp.Network.Carriers[0]
	}

	loc := resp.Location
	sec := resp.HazardReport

	hasThreat := false
	switch {
	case sec.IsSpamhausDrop, sec.IsSpamhausEdrop,
		sec.IsSpamhausAsnDrop, sec.IsBlacklistedUceprotect,
		sec.IsBlacklistedBlocklistDe, sec.IsUnreachable:
		hasThreat = true
	}
	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             resp.IP,
		Organization:   resp.Network.Organisation,
		ASNumber:       carrier.ASNNumeric,
		IsHosting:      sec.IsHostingASN || sec.HostingLikelihood >= 3,
		IsProxy:        sec.IsKnownAsProxy,
		IsAnonymousVPN: sec.IsKnownAsVPN,
		IsTor:          sec.IsKnownAsTorServer,
		IsBogon:        sec.IsBogon,
		HasOtherThreat: hasThreat,
		Country:        resp.Country.ISOAlpha2,
		City:           loc.City,
		Latitude:       loc.Latitude,
		Longitude:      loc.Longitude,
	}, nil
}

func (p BigDataCloudProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.IPGeolocationFull(ipaddr)
}
