package provider

import "errors"

type FraudCheckResponse struct {
	ServiceName  string `json:"service_name"`
	IP           string `json:"ip"`
	Hostname     string `json:"hostname"`
	ISP          string `json:"isp"`
	Organization string `json:"organization"`
	ASNumber     int64  `json:"asn"`

	RiskScore      float64 `json:"risk_score"` // 0 ~ 1
	IsVPN          bool    `json:"is_vpn"`
	IsHosting      bool    `json:"is_hosting"`
	IsProxy        bool    `json:"is_proxy"`
	IsTor          bool    `json:"is_tor"`
	IsBot          bool    `json:"is_bot"`
	IsBogon        bool    `json:"is_bogon"`
	HasOtherThreat bool    `json:"has_other_threat"`
	ThreatComment  string  `json:"threat_comment"`

	Country   string  `json:"country"`
	City      string  `json:"city"`
	Region    string  `json:"region"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`

	Err string `json:"error"`
}

func (r FraudCheckResponse) HasError() bool {
	return r.Err != ""
}

func (r FraudCheckResponse) Error() error {
	if !r.HasError() {
		return nil
	}
	return errors.New(r.Err)
}
