package provider

// Provider is a interface for IP address check service provider.
type Provider interface {
	Init(conf Config) error
	String() string
	CheckIP(ipaddr string) (FraudCheckResponse, error)
	RawCheckIP(ipaddr string) (interface{}, error)
}

type Config interface{}
