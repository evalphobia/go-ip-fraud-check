package abuseipdb

type Response struct {
	Data Data `json:"data"`
}

type Data struct {
	IPAddress            string   `json:"ipAddress"`
	IPVersion            int64    `json:"ipVersion"`
	ISP                  string   `json:"isp"`
	Domain               string   `json:"domain"`
	Hostnames            []string `json:"hostnames"`
	CountryCode          string   `json:"countryCode"`
	CountryName          string   `json:"countryName"`
	IsPublic             bool     `json:"isPublic"`
	IsWhitelisted        bool     `json:"isWhitelisted"`
	AbuseConfidenceScore int64    `json:"abuseConfidenceScore"`
	UsageType            string   `json:"usageType"`
	TotalReports         int64    `json:"totalReports"`
	NumDistinctUsers     int64    `json:"numDistinctUsers"`
	LastReportedAt       string   `json:"lastReportedAt"`
	Reports              []Report `json:"reports"`
}

func (d Data) IsHosting() bool {
	_, ok := hostingMap[d.UsageType]
	return ok
}

type Report struct {
	ReportedAt          string  `json:"reportedAt"`
	Comment             string  `json:"comment"`
	Categories          []int64 `json:"categories"`
	ReporterID          int64   `json:"reporterId"`
	ReporterCountryCode string  `json:"reporterCountryCode"`
	ReporterCountryName string  `json:"reporterCountryName"`
}

var hostingMap = map[string]struct{}{
	"Data Center/Web Hosting/Transit": struct{}{},
}
