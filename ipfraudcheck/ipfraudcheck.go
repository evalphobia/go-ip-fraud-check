package ipfraudcheck

import (
	"errors"

	"github.com/evalphobia/go-ip-fraud-check/log"
	"github.com/evalphobia/go-ip-fraud-check/provider"
)

type Client struct {
	logger    log.Logger
	providers []provider.Provider
	useRoutes bool
}

func New(conf Config, providers []provider.Provider) (*Client, error) {
	enabledProviders := make([]provider.Provider, 0, len(providers))
	for _, p := range providers {
		err := p.Init(conf)
		if err != nil {
			return nil, err
		}
		enabledProviders = append(enabledProviders, p)
	}

	if len(enabledProviders) == 0 {
		return nil, errors.New("no ip check providers are specified")
	}

	logger := conf.GetLogger()

	for _, e := range enabledProviders {
		logger.Infof("Use %s\n", e.String())
	}

	return &Client{
		logger:    logger,
		providers: enabledProviders,
		useRoutes: conf.UseRoute,
	}, nil
}

func (c Client) RawCheckIP(ipaddr string) ([]interface{}, error) {
	list := make([]interface{}, len(c.providers))
	for i, p := range c.providers {
		resp, err := p.RawCheckIP(ipaddr)
		if err != nil {
			return nil, err
		}
		list[i] = resp
	}

	return list, nil
}

func (c Client) CheckIP(ipaddr string) (Response, error) {
	list := make([]provider.FraudCheckResponse, len(c.providers))
	for i, p := range c.providers {
		resp, err := p.CheckIP(ipaddr)
		if err != nil {
			resp.Err = err.Error()
		}
		list[i] = resp
	}

	resp := Response{
		List: list,
	}
	if c.useRoutes {
		asn, ok := resp.FindASN()
		if ok {
			cli := NewWhoisClient()
			routes, err := cli.GetRoutes(asn)
			if err != nil {
				return resp, err
			}
			resp.ASPrefix = routes
		}
	}
	return resp, nil
}

type Response struct {
	List     []provider.FraudCheckResponse `json:"list"`
	ASPrefix []string                      `json:"as_prefix"`
}

func (r Response) FindASN() (int64, bool) {
	for _, resp := range r.List {
		if resp.ASNumber != 0 {
			return resp.ASNumber, true
		}
	}
	return 0, false
}
