package spur

import (
	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://api.spur.us"
)

type Client struct {
	privateclient.RESTClient
}

func NewClient(token string) Client {
	return Client{
		RESTClient: privateclient.RESTClient{
			Option: privateclient.Option{
				Headers: map[string]string{"Token": token},
				BaseURL: defaultBaseURL,
			},
		},
	}
}

func (c *Client) SetDebug(b bool) {
	c.RESTClient.Debug = b
}

func (c Client) DoContext(ipaddr string) (Response, error) {
	resp := Response{}
	err := c.RESTClient.CallGET("/v1/context/"+ipaddr, nil, &resp)
	return resp, err
}
