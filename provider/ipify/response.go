package ipify

type Response struct {
	IP      string   `json:"ip"`
	ISP     string   `json:"isp"`
	Domains []string `json:"domains"`

	Type          string   `json:"type"`
	ContinentCode string   `json:"continent_code"`
	ContinentName string   `json:"continent_name"`
	CountryCode   string   `json:"country_code"`
	CountryName   string   `json:"country_name"`
	RegionCode    string   `json:"region_code"`
	RegionName    string   `json:"region_name"`
	City          string   `json:"city"`
	Zip           string   `json:"zip"`
	AS            AS       `json:"as"`
	Location      Location `json:"location"`
	Proxy         Proxy    `json:"proxy"`
}

type AS struct {
	ASN    int64  `json:"asn"`
	Name   string `json:"name"`
	Route  string `json:"route"`
	Domain string `json:"domain"`
	Type   string `json:"type"`
}

type Location struct {
	GeonameID  int64   `json:"geoname_id"`
	Country    string  `json:"country"`
	Region     string  `json:"region"`
	City       string  `json:"city"`
	Latitude   float64 `json:"lat"`
	Longitude  float64 `json:"lng"`
	PostalCode string  `json:"postalCode"`
	Timezone   string  `json:"timezone"`
}

type Proxy struct {
	Proxy bool `json:"proxy"`
	VPN   bool `json:"vpn"`
	Tor   bool `json:"tor"`
}
