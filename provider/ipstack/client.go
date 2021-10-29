package ipstack

import (
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://api.ipstack.com"
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
		Security: true,
		Fields: []string{
			"main",
			"connection",
			"security",
		},
	})
}

func (c Client) LookUp(ipaddr string, opt Option) (Response, error) {
	params := make(map[string]string)
	params["access_key"] = c.apiKey
	if opt.Hostname {
		params["hostname"] = "1"
	}
	if opt.Security {
		params["security"] = "1"
	}
	if len(opt.Fields) != 0 {
		params["fields"] = opt.GetFields()
	}
	if opt.Language != "" {
		params["language"] = opt.Language
	}
	if opt.Callback != "" {
		params["callback"] = opt.Callback
	}
	if opt.Output != "" {
		params["output"] = opt.Output
	}

	resp := Response{}
	err := c.RESTClient.CallGET("/"+ipaddr, params, &resp)
	return resp, err
}

type Option struct {
	Fields   []string
	Hostname bool
	Security bool
	Language string
	Callback string
	Output   string
}

func (o Option) GetFields() string {
	return strings.Join(o.Fields, ",")
}
