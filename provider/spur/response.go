package spur

type Response struct {
	IP              string          `json:"ip"`
	Anonymous       bool            `json:"anonymous"`
	AS              AS              `json:"as"`
	Assignment      Assignment      `json:"assignment"`
	DeviceBehaviors DeviceBehaviors `json:"deviceBehaviors"`
	Devices         Devices         `json:"devices"`
	GeoLite         GeoLite         `json:"geoLite"`
	GeoPrecision    GeoPrecision    `json:"geoPrecision"`
	Infrastructure  string          `json:"infrastructure"`
	ProxiedTraffic  ProxiedTraffic  `json:"proxiedTraffic"`
	SimilarIPs      SimilarIPs      `json:"similarIPs"`
	VPNOperators    VPNOperators    `json:"VPNOperators"`
	WiFi            WiFi            `json:"wifi"`
}

func (r Response) IsProxy() bool {
	return r.ProxiedTraffic.Exists
}

func (r Response) IsVPN() bool {
	return r.VPNOperators.Exists
}

func (r Response) IsHosting() bool {
	return r.Infrastructure == "DATACENTER"
}

func (r Response) IsTor() bool {
	for _, v := range r.DeviceBehaviors.Behaviors {
		if v.Name == "TOR_PROXY_USER" {
			return true
		}
	}
	return false
}

func (r Response) IsMobile() bool {
	return r.Infrastructure == "MOBILE"
}

func (r Response) IsFileSharing() bool {
	for _, v := range r.DeviceBehaviors.Behaviors {
		if v.Name == "FILE_SHARING" {
			return true
		}
	}
	return false
}

type AS struct {
	Number       int64  `json:"number"`
	Organization string `json:"organization"`
}

type Assignment struct {
	Exists       bool   `json:"exists"`
	LastTurnover string `json:"lastTurnover"`
}

type DeviceBehaviors struct {
	Exists    bool       `json:"exists"`
	Behaviors []Behavior `json:"behaviors"`
}

type Behavior struct {
	Name string `json:"name"`
}

type Devices struct {
	Estimate int64 `json:"estimate"`
}

type GeoLite struct {
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
}

type GeoPrecision struct {
	Exists  bool   `json:"exists"`
	City    string `json:"city"`
	Country string `json:"country"`
	State   string `json:"state"`
	Hash    string `json:"hash"`
	Spread  string `json:"spread"`
	Point   Point  `json:"point"`
}

type Point struct {
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Radius    int64   `json:"radius"`
}

type ProxiedTraffic struct {
	Exists  bool    `json:"exists"`
	Proxies []Proxy `json:"proxies"`
}

type Proxy struct {
	Name string `json:"name"`
	Type string `json:"type"`
}

type SimilarIPs struct {
	Exists bool     `json:"exists"`
	IPs    []string `json:"ips"`
}

type VPNOperators struct {
	Exists    bool       `json:"exists"`
	Operators []Operator `json:"operators"`
}

type Operator struct {
	Name string `json:"name"`
}

type WiFi struct {
	Exists bool `json:"exists"`
}
