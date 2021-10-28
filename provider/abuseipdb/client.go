package abuseipdb

import (
	"github.com/evalphobia/go-ip-fraud-check/provider/privateclient"
)

const (
	defaultBaseURL = "https://api.abuseipdb.com"
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
				Headers: map[string]string{"KEY": apiKey},
				BaseURL: defaultBaseURL,
			},
		},
	}
}

func (c *Client) SetDebug(b bool) {
	c.RESTClient.Debug = b
}

func (c Client) Check(ipaddr string) (Response, error) {
	params := make(map[string]string)
	params["ipAddress"] = ipaddr
	params["maxAgeInDays"] = "365"
	params["verbose"] = ""

	resp := Response{}
	err := c.RESTClient.CallGET("/api/v2/check", params, &resp)
	return resp, err
}
