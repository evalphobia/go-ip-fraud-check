package ipstack

type Response struct {
	IP            string     `json:"ip"`
	Hostname      string     `json:"hostname"`
	Type          string     `json:"type"`
	ContinentCode string     `json:"continent_code"`
	ContinentName string     `json:"continent_name"`
	CountryCode   string     `json:"country_code"`
	CountryName   string     `json:"country_name"`
	RegionCode    string     `json:"region_code"`
	RegionName    string     `json:"region_name"`
	City          string     `json:"city"`
	Zip           string     `json:"zip"`
	Latitude      float64    `json:"latitude"`
	Longitude     float64    `json:"longitude"`
	Location      Location   `json:"location"`
	Timezone      Timezone   `json:"timezone"`
	Currency      Currency   `json:"currency"`
	Connection    Connection `json:"connection"`
	Security      Security   `json:"security"`
}

type Location struct {
	GeonameID               int64      `json:"geoname_id"`
	Capital                 string     `json:"capital"`
	CountryFlag             string     `json:"country_flag"`
	CountryFlagEmoji        string     `json:"country_flag_emoji"`
	CountryFlagEmojiUnicode string     `json:"country_flag_emoji_unicode"`
	CallingCode             string     `json:"calling_code"`
	IsEU                    bool       `json:"is_eu"`
	Languages               []Language `json:"languages"`
}

type Language struct {
	Code   string `json:"code"`
	Name   string `json:"name"`
	Native string `json:"native"`
}

type Timezone struct {
	ID               string `json:"id"`
	CurrentTime      string `json:"current_time"`
	GMTOffset        int64  `json:"gmt_offset"`
	Code             string `json:"code"`
	IsDaylightSaving bool   `json:"is_daylight_saving"`
}

type Currency struct {
	Code         string `json:"code"`
	Name         string `json:"name"`
	Plural       string `json:"plural"`
	Symbol       string `json:"symbol"`
	SymbolNative string `json:"symbol_native"`
}

type Connection struct {
	ASN int64  `json:"asn"`
	ISP string `json:"isp"`
}

type Security struct {
	IsProxy     bool     `json:"is_proxy"`
	ProxyType   string   `json:"proxy_type"`
	IsCrawler   bool     `json:"is_crawler"`
	CrawlerName string   `json:"crawler_name"`
	CrawlerType string   `json:"crawler_type"`
	IsTor       bool     `json:"is_tor"`
	ThreatLevel string   `json:"threat_level"`
	ThreatTypes []string `json:"threat_types"`
}
