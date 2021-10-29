package dbip

import (
	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://api.db-ip.com"
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

func (c Client) LookUp(ipaddr string) (Response, error) {
	params := make(map[string]string)

	resp := Response{}
	err := c.RESTClient.CallGET("/v2/"+c.apiKey+"/"+ipaddr, params, &resp)
	return resp, err
}
