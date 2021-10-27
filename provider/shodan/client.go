package shodan

import (
	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://api.shodan.io"
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

func (c Client) GetHost(ipaddr string, param ...HostParameter) (HostResponse, error) {
	params := make(map[string]interface{})
	params["key"] = c.apiKey

	var p HostParameter
	if len(param) != 0 {
		p = param[0]
	}
	if p.History {
		params["history"] = true
	}
	if p.Minify {
		params["minify"] = true
	}

	resp := HostResponse{}
	err := c.RESTClient.CallGET("/shodan/host/"+ipaddr, params, &resp)
	return resp, err
}

type HostParameter struct {
	History bool
	Minify  bool
}
