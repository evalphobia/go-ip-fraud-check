package ipgeolocation

import (
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://api.ipgeolocation.io/ipgeo"
)

type Client struct {
	apiKey string
	privateclient.RESTClient
}

func NewClient(apiKey string) Client {
	return Client{
		apiKey: apiKey,
		RESTClient: privateclient.RESTClient{
			Option: privateclient.Option{
				BaseURL: defaultBaseURL,
			},
		},
	}
}

func (c *Client) SetDebug(b bool) {
	c.RESTClient.Debug = b
}

func (c Client) LookUpWithSecurity(ipaddr string) (Response, error) {
	return c.CallGET(ipaddr, RequestParameter{
		IncludeSecurity: true,
	})
}

func (c Client) CallGET(ipaddr string, p RequestParameter) (Response, error) {
	params := make(map[string]string)
	params["apiKey"] = c.apiKey
	params["ip"] = ipaddr

	if p.Language != "" {
		params["lang"] = p.Language
	}
	if len(p.Fields) != 0 {
		params["fields"] = strings.Join(p.Fields, ",")
	}
	if len(p.ExcludeFields) != 0 {
		params["excludes"] = strings.Join(p.ExcludeFields, ",")
	}

	var includes []string
	if p.IncludeSecurity {
		includes = append(includes, "security")
	}
	if p.IncludeUserAgent {
		includes = append(includes, "useragent")
	}
	if p.IncludeHostname {
		includes = append(includes, "hostname")
	}
	if p.IncludeLiveHostname {
		includes = append(includes, "liveHostname")
	}
	if p.IncludeHostnameFallbackLive {
		includes = append(includes, "hostnameFallbackLive")
	}
	if len(includes) != 0 {
		params["include"] = strings.Join(includes, ",")
	}

	resp := Response{}
	err := c.RESTClient.CallGET("", params, &resp)
	return resp, err
}

type RequestParameter struct {
	Language                    string
	Fields                      []string
	ExcludeFields               []string
	IncludeSecurity             bool
	IncludeUserAgent            bool
	IncludeHostname             bool
	IncludeLiveHostname         bool
	IncludeHostnameFallbackLive bool
}
