package ipapicom

import (
	"strconv"
	"strings"
)

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`

	Query         string  `json:"query"`
	ISP           string  `json:"isp"`
	Org           string  `json:"org"`
	AS            string  `json:"as"`
	ASName        string  `json:"asname"`
	Reverse       string  `json:"reverse"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionCode    string  `json:"regionCode"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Latitude      float64 `json:"lat"`
	Longitude     float64 `json:"lng"`
	Timezone      string  `json:"timezone"`
	Offset        int64   `json:"offset"`
	Curency       int64   `json:"curency"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}

func (r Response) ASNumber() int64 {
	if r.AS == "" {
		return 0
	}
	parts := strings.Split(r.AS, " ")
	asn, _ := strconv.ParseInt(strings.TrimPrefix(parts[0], "AS"), 10, 64)
	return asn
}
