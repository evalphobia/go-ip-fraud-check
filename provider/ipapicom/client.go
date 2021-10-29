package ipapicom

import (
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "http://ip-api.com/json/"
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
	return c.LookUp(ipaddr, Option{
		Fields: []string{
			"status",
			"message",
			"countryCode",
			"regionName",
			"city",
			"lat",
			"lon",
			"isp",
			"org",
			"as",
			"proxy",
			"hosting",
		},
	})
}

func (c Client) LookUp(ipaddr string, opt Option) (Response, error) {
	params := make(map[string]string)

	if len(opt.Fields) != 0 {
		params["fields"] = opt.GetFields()
	}
	if opt.Lang != "" {
		params["lang"] = opt.Lang
	}
	if opt.Callback != "" {
		params["callback"] = opt.Callback
	}

	resp := Response{}
	err := c.RESTClient.CallGET("/"+ipaddr, params, &resp)
	return resp, err
}

type Option struct {
	Fields   []string
	Lang     string
	Callback string
}

func (o Option) GetFields() string {
	return strings.Join(o.Fields, ",")
}
