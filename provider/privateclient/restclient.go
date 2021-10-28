package privateclient

import (
	"fmt"

	"github.com/evalphobia/httpwrapper/request"
)

// RESTClient is http REST client.
type RESTClient struct {
	Option
	BasicAuthUser string
	BasicAuthPass string
}

func New() *RESTClient {
	return &RESTClient{}
}

func (c *RESTClient) SetAuthData(user, pass string) {
	c.BasicAuthUser = user
	c.BasicAuthPass = pass
}

func (c *RESTClient) SetOption(opt Option) {
	c.Option = opt
}

// CallGET sends GET request to `url` with `params` and set reqponse to `result`.
func (c *RESTClient) CallGET(path string, params, result interface{}) (err error) {
	opt := c.Option
	url := fmt.Sprintf("%s%s", opt.BaseURL, path)

	resp, err := request.GET(url, request.Option{
		Query:     params,
		User:      c.BasicAuthUser,
		Pass:      c.BasicAuthPass,
		Retry:     opt.Retry,
		Debug:     opt.Debug,
		Headers:   opt.getHeaders(),
		UserAgent: opt.getUserAgent(),
		Timeout:   opt.getTimeout(),
	})
	if err != nil {
		return err
	}
	if err := resp.HasStatusCodeError(); err != nil {
		return err
	}
	return resp.JSON(result)
}

// CallPOST sends POST request to `url` with `params` and set reqponse to `result`.
func (c *RESTClient) CallPOST(path string, params, result interface{}) (err error) {
	opt := c.Option
	url := fmt.Sprintf("%s%s", opt.BaseURL, path)

	resp, err := request.POST(url, request.Option{
		Payload:     params,
		PayloadType: request.PayloadTypeJSON,
		User:        c.BasicAuthUser,
		Pass:        c.BasicAuthPass,
		Retry:       opt.Retry,
		Debug:       opt.Debug,
		Headers:     opt.getHeaders(),
		UserAgent:   opt.getUserAgent(),
		Timeout:     opt.getTimeout(),
	})
	if err != nil {
		return err
	}
	if err := resp.HasStatusCodeError(); err != nil {
		return err
	}
	return resp.JSON(result)
}
