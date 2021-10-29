package ipify

import (
	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://geo.ipify.org"
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
	return c.LookUp("country,city,vpn", Option{
		IPAddress: ipaddr,
	})
}

func (c Client) LookUpWithCity(ipaddr string) (Response, error) {
	return c.LookUp("country,city", Option{
		IPAddress: ipaddr,
	})
}

func (c Client) LookUpCountry(ipaddr string) (Response, error) {
	return c.LookUp("country", Option{
		IPAddress: ipaddr,
	})
}

func (c Client) LookUp(typ string, opt Option) (Response, error) {
	params := make(map[string]string)
	params["apiKey"] = c.apiKey

	if opt.IPAddress != "" {
		params["ipAddress"] = opt.IPAddress
	}
	if opt.Domain != "" {
		params["domain"] = opt.Domain
	}
	if opt.Email != "" {
		params["email"] = opt.Email
	}
	if opt.EscapedUnicode {
		params["escapedUnicode"] = "1"
	}

	resp := Response{}
	err := c.RESTClient.CallGET("/api/v2/"+typ, params, &resp)
	return resp, err
}

type Option struct {
	IPAddress      string
	Domain         string
	Email          string
	EscapedUnicode bool
}
