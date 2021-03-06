package ipfraudcheck

import (
	"os"
	"reflect"
	"time"

	"github.com/evalphobia/go-ip-fraud-check/log"
)

const (
	envAbuseIPDBAPIKey     = "FRAUD_CHECK_ABUSEIPDB_APIKEY"
	envDBIPAPIKey          = "FRAUD_CHECK_DBIP_APIKEY"
	envIP2ProxyAPIKey      = "FRAUD_CHECK_IP2PROXY_APIKEY"
	envIP2ProxyAPIPackage  = "FRAUD_CHECK_IP2PROXY_PACKAGE"
	envIPAPIComAPIKey      = "FRAUD_CHECK_IPAPICOM_APIKEY"
	envIPdatacoAPIKey      = "FRAUD_CHECK_IPDATACO_APIKEY"
	envIPGeolocationAPIKey = "FRAUD_CHECK_IPGEOLOCATION_APIKEY"
	envIPifyAPIKey         = "FRAUD_CHECK_IPIFY_APIKEY"
	envIPinfoioToken       = "FRAUD_CHECK_IPINFOIO_TOKEN"
	envIPStackAPIKey       = "FRAUD_CHECK_IPSTACK_APIKEY"
	envShodanAPIKey        = "FRAUD_CHECK_SHODAN_APIKEY"
	envSpurToken           = "FRAUD_CHECK_SPUR_TOKEN"
)

const (
	// ref: https://www.ip2location.com/web-service/ip2proxy
	defaultIP2ProxyAPIPackage = "PX2"
)

// Config contains parameters for IP check API providers.
type Config struct {
	// abuseipdb.com
	AbuseIPDBAPIKey string
	// db-ip.com
	DBIPAPIKey string
	// bigdatacloud.com
	BigDataCloudAPIKey string
	// ip2location.com
	IP2ProxyAPIKey     string
	IP2ProxyAPIPackage string
	// ip-api.com
	IPAPIComAPIKey string
	// ipdata.co
	IPdatacoAPIKey string
	// ipgeolocation.io
	IPGeolocationAPIKey string
	// ipify.org
	IPifyAPIKey string
	// ipinfo.io
	IPinfoioToken string
	// ipqualityscore.com
	IPQualityScoreAPIKey string
	// ipregistry.co
	IPRegistryAPIKey string
	// ipstack.com
	IPStackAPIKey string
	// minFraud
	MinFraudAccountID  string
	MinFraudLicenseKey string
	// shodan.io
	ShodanAPIKey string
	// spur.us
	SpurToken string

	// common option
	UseRoute bool
	Interval time.Duration
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

func (c Config) GetAbuseIPDBAPIKey() string {
	s := os.Getenv(envAbuseIPDBAPIKey)
	if s != "" {
		return s
	}
	return c.AbuseIPDBAPIKey
}

func (c Config) GetDBIPAPIKey() string {
	s := os.Getenv(envDBIPAPIKey)
	if s != "" {
		return s
	}
	return c.DBIPAPIKey
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

func (c Config) GetIPAPIComAPIKey() string {
	s := os.Getenv(envIPAPIComAPIKey)
	if s != "" {
		return s
	}
	return c.IPAPIComAPIKey
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

func (c Config) GetIPifyAPIKey() string {
	s := os.Getenv(envIPifyAPIKey)
	if s != "" {
		return s
	}
	return c.IPifyAPIKey
}

func (c Config) GetIPinfoioToken() string {
	s := os.Getenv(envIPinfoioToken)
	if s != "" {
		return s
	}
	return c.IPinfoioToken
}

func (c Config) GetIPStackAPIKey() string {
	s := os.Getenv(envIPStackAPIKey)
	if s != "" {
		return s
	}
	return c.IPStackAPIKey
}

func (c Config) GetShodanAPIKey() string {
	s := os.Getenv(envShodanAPIKey)
	if s != "" {
		return s
	}
	return c.ShodanAPIKey
}

func (c Config) GetSpurToken() string {
	s := os.Getenv(envSpurToken)
	if s != "" {
		return s
	}
	return c.SpurToken
}
