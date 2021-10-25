package ipgeolocation

import (
	"errors"
	"strconv"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type IPGeoLocationProvider struct {
	client Client
}

func (p *IPGeoLocationProvider) Init(conf provider.Config) error {
	c, ok := conf.(ipfraudcheck.Config)
	if !ok {
		return errors.New("incompatible config type for IPGeoLocationProvider")
	}

	apiKey := c.GetIPGeolocationAPIKey()
	if apiKey == "" {
		return errors.New("apikey for IPGeolocation is empty. you must set directly or use 'FRAUD_CHECK_IPGEOLOCATION_APIKEY' envvar")
	}
	cli := NewClient(apiKey)
	cli.SetDebug(c.Debug)
	p.client = cli
	return nil
}

func (p IPGeoLocationProvider) String() string {
	return "IPGeolocation"
}

func (p IPGeoLocationProvider) CheckIP(ipaddr string) (provider.FraudCheckResponse, error) {
	emptyResult := provider.FraudCheckResponse{}

	resp, err := p.client.LookUpWithSecurity(ipaddr)
	if err != nil {
		return emptyResult, err
	}

	var lat, longi float64
	if resp.Latitude != "" {
		lat, _ = strconv.ParseFloat(resp.Latitude, 64)
	}
	if resp.Longitude != "" {
		longi, _ = strconv.ParseFloat(resp.Longitude, 64)
	}

	asn, _ := strconv.ParseInt(resp.ASN, 10, 64)
	sec := resp.Security
	return provider.FraudCheckResponse{
		ServiceName:    p.String(),
		IP:             ipaddr,
		Hostname:       resp.Hostname,
		ISP:            resp.ISP,
		ASNumber:       asn,
		Organization:   resp.Organization,
		RiskScore:      float64(sec.ThreatScore) / 100,
		IsProxy:        sec.IsProxy,
		IsAnonymous:    sec.IsAnonymous,
		IsAnonymousVPN: sec.ProxyType == "VPN",
		IsHosting:      sec.IsCloudProvider,
		IsTor:          sec.IsTor,
		HasOtherThreat: sec.IsKnownAttacker,
		Country:        resp.CountryCode2,
		City:           resp.City,
		Latitude:       lat,
		Longitude:      longi,
	}, nil
}

func (p IPGeoLocationProvider) RawCheckIP(ipaddr string) (interface{}, error) {
	return p.client.LookUpWithSecurity(ipaddr)
}
