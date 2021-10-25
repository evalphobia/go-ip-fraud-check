package ipfraudcheck

import (
	"os"
	"reflect"
	"time"

	"github.com/evalphobia/go-ip-fraud-check/log"
)

const (
	envIP2ProxyAPIKey      = "FRAUD_CHECK_IP2PROXY_APIKEY"
	envIP2ProxyAPIPackage  = "FRAUD_CHECK_IP2PROXY_PACKAGE"
	envIPdatacoAPIKey      = "FRAUD_CHECK_IPDATACO_APIKEY"
	envIPGeolocationAPIKey = "FRAUD_CHECK_IPGEOLOCATION_APIKEY"
	envIPinfoioToken       = "FRAUD_CHECK_IPINFOIO_TOKEN"
)

const (
	// ref: https://www.ip2location.com/web-service/ip2proxy
	defaultIP2ProxyAPIPackage = "PX2"
)

// Config contains parameters for IP check API providers.
type Config struct {
	// ip2location.com
	IP2ProxyAPIKey     string
	IP2ProxyAPIPackage string
	// ipdata.co
	IPdatacoAPIKey string
	// ipgeolocation.io
	IPGeolocationAPIKey string
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

func (c Config) GetIP2ProxyAPIKey() string {
	s := os.Getenv(envIP2ProxyAPIKey)
	if s != "" {
		return s
	}
	return c.IP2ProxyAPIKey
}

func (c Config) GetIP2ProxyAPIPackage() string {
	s := os.Getenv(envIP2ProxyAPIPackage)
	if s != "" {
		return s
	}
	if c.IP2ProxyAPIPackage != "" {
		return c.IP2ProxyAPIPackage
	}
	return defaultIP2ProxyAPIPackage
}

func (c Config) GetIPdatacoAPIKey() string {
	s := os.Getenv(envIPdatacoAPIKey)
	if s != "" {
		return s
	}
	return c.IPdatacoAPIKey
}

func (c Config) GetIPGeolocationAPIKey() string {
	s := os.Getenv(envIPGeolocationAPIKey)
	if s != "" {
		return s
	}
	return c.IPGeolocationAPIKey
}

func (c Config) GetIPinfoioToken() string {
	s := os.Getenv(envIPinfoioToken)
	if s != "" {
		return s
	}
	return c.IPinfoioToken
}
