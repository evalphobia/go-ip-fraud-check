package ipfraudcheck

import (
	"os"
	"reflect"
	"time"

	"github.com/evalphobia/go-ip-fraud-check/log"
)

const (
	envIPdatacoAPIKey = "FRAUD_CHECK_IPDATACO_APIKEY"
	envIPinfoioToken  = "FRAUD_CHECK_IPINFOIO_TOKEN"
)

// Config contains parameters for IP check API providers.
type Config struct {
	// ipdata.con
	IPdatacoAPIKey string
	// ipinfo.io
	IPinfoioToken string
	// minFraud
	MinFraudAccountID  string
	MinFraudLicenseKey string

	// common option
	UseRoute bool
	Debug    bool
	Timeout  time.Duration
	Logger   log.Logger
}

func (c Config) GetLogger() log.Logger {
	switch {
	case c.Logger == nil,
		reflect.ValueOf(c.Logger).IsNil():
		return log.DefaultLogger
	}
	return c.Logger
}

func (c Config) GetIPdatacoAPIKey() string {
	s := os.Getenv(envIPdatacoAPIKey)
	if s != "" {
		return s
	}
	return c.IPdatacoAPIKey
}

func (c Config) GetIPinfoioToken() string {
	s := os.Getenv(envIPinfoioToken)
	if s != "" {
		return s
	}
	return c.IPinfoioToken
}
