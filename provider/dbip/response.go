package dbip

type Response struct {
	IPAddress     string   `json:"ipAddress"`
	ContinentCode string   `json:"continentCode"`
	ContinentName string   `json:"continentName"`
	CountryCode   string   `json:"countryCode"`
	CountryName   string   `json:"countryName"`
	IsEuMember    bool     `json:"isEuMember"`
	CurrencyCode  string   `json:"currencyCode"`
	CurrencyName  string   `json:"currencyName"`
	PhonePrefix   string   `json:"phonePrefix"`
	Languages     []string `json:"languages"`
	StateProvCode string   `json:"stateProvCode"`
	StateProv     string   `json:"stateProv"`
	District      string   `json:"district"`
	City          string   `json:"city"`
	GeonameID     int64    `json:"geonameId"`
	ZipCode       string   `json:"zipCode"`
	Latitude      float64  `json:"latitude"`
	Longitude     float64  `json:"longitude"`
	GMTOffset     int64    `json:"gmtOffset"`
	TimeZone      string   `json:"timeZone"`
	WeatherCode   string   `json:"weatherCode"`
	ASNumber      int64    `json:"asNumber"`
	ASName        string   `json:"asName"`
	ISP           string   `json:"isp"`
	LinkType      string   `json:"linkType"`
	UsageType     string   `json:"usageType"`
	Organization  string   `json:"organization"`
	IsCrawler     bool     `json:"isCrawler"`
	CrawlerName   string   `json:"crawlerName"`
	IsProxy       bool     `json:"isProxy"`
	ProxyType     string   `json:"proxyType"`
	ThreatLevel   string   `json:"threatLevel"`
	ThreatDetails []string `json:"threatDetails"`
}

func (r Response) IsHTTPProxy() bool {
	return r.ProxyType == "http"
}

func (r Response) IsVPN() bool {
	return r.ProxyType == "vpn"
}

func (r Response) IsTor() bool {
	return r.ProxyType == "tor"
}

func (r Response) IsHosting() bool {
	return r.UsageType == "hosting"
}

func (r Response) IsBot() bool {
	for _, v := range r.ThreatDetails {
		if v == "bot" {
			return true
		}
	}
	return false
}

func (r Response) HasOtherThreat() bool {
	for _, v := range r.ThreatDetails {
		if v == "attack-source" {
			return true
		}
	}
	return false
}

func (r Response) GetRiskScore() float64 {
	switch r.ThreatLevel {
	case "medium":
		return 0.5
	case "high":
		return 0.75
	default:
		return 0
	}
}
