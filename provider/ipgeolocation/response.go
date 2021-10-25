package ipgeolocation

type Response struct {
	IP             string    `json:"ip"`
	Hostname       string    `json:"hostname"`
	ContinentCode  string    `json:"continent_code"`
	ContinentName  string    `json:"continent_name"`
	CountryCode2   string    `json:"country_code2"`
	CountryCode3   string    `json:"country_code3"`
	CountryName    string    `json:"country_name"`
	CountryCapital string    `json:"country_capital"`
	StateProv      string    `json:"state_prov"`
	District       string    `json:"district"`
	City           string    `json:"city"`
	ZipCode        string    `json:"zipcode"`
	Latitude       string    `json:"latitude"`
	Longitude      string    `json:"longitude"`
	IsEU           bool      `json:"is_eu"`
	CallingCode    string    `json:"calling_code"`
	CountryTLD     string    `json:"country_tld"`
	Languages      string    `json:"languages"`
	CountryFlag    string    `json:"country_flag"`
	GeonameID      string    `json:"geoname_id"`
	ISP            string    `json:"isp"`
	ConnectionType string    `json:"connection_type"`
	Organization   string    `json:"organization"`
	ASN            string    `json:"asn"`
	Currency       Currency  `json:"currency"`
	TimeZone       TimeZone  `json:"time_zone"`
	Security       Security  `json:"security"`
	UserAgent      UserAgent `json:"user_agent"`
}

type Currency struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Symbol string `json:"symbol"`
}

type TimeZone struct {
	Name            string  `json:"name"`
	Offset          int     `json:"offset"`
	CurrentTime     string  `json:"current_time"`
	CurrentTimeUnix float64 `json:"current_time_unix"`
	IsDST           bool    `json:"is_dst"`
	DSTSavings      int     `json:"dst_savings"`
}

type Security struct {
	ThreatScore     int    `json:"threat_score"`
	IsTor           bool   `json:"is_tor"`
	IsProxy         bool   `json:"is_proxy"`
	ProxyType       string `json:"proxy_type"`
	IsAnonymous     bool   `json:"is_anonymous"`
	IsKnownAttacker bool   `json:"is_known_attacker"`
	IsCloudProvider bool   `json:"is_cloud_provider"`
}

type UserAgent struct {
	UserAgentString string          `json:"userAgentString"`
	Name            string          `json:"name"`
	Type            string          `json:"type"`
	Version         string          `json:"version"`
	VersionMajor    string          `json:"versionMajor"`
	Device          Device          `json:"device"`
	Engine          Engine          `json:"engine"`
	OperatingSystem OperatingSystem `json:"operatingSystem"`
}

type Device struct {
	Name  string `json:"name"`
	Type  string `json:"type"`
	Brand string `json:"brand"`
	CPU   string `json:"CPU"`
}

type Engine struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	VersionMajor string `json:"versionMajor"`
}

type OperatingSystem struct {
	Name         string `json:"name"`
	Type         string `json:"type"`
	Version      string `json:"version"`
	VersionMajor string `json:"versionMajor"`
}
