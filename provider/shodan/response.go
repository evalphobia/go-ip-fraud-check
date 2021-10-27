package shodan

type HostResponse struct {
	RegionCode   string     `json:"region_code"`
	IP           int64      `json:"ip"`
	IPStr        string     `json:"ip_str"`
	Org          string     `json:"org"`
	ISP          string     `json:"isp"`
	ASN          string     `json:"asn"`
	AreaCode     string     `json:"area_code"`
	PostalCode   string     `json:"postal_code"`
	CountryCode  string     `json:"country_code"`
	CountryCode3 string     `json:"country_code3"`
	CountryName  string     `json:"country_name"`
	City         string     `json:"city"`
	DMACode      string     `json:"dma_code"`
	LastUpdate   string     `json:"last_update"`
	Latitude     float64    `json:"latitude"`
	Longitude    float64    `json:"longitude"`
	Hostnames    []string   `json:"hostnames"`
	Domains      []string   `json:"domains"`
	Ports        []int      `json:"ports"`
	Tags         []string   `json:"tags"`
	OS           string     `json:"os"`
	Data         []HostData `json:"data"`
}

type HostData struct {
	Hash      int64    `json:"hash"`
	Timestamp string   `json:"timestamp"`
	IP        int64    `json:"ip"`
	IPStr     string   `json:"ip_str"`
	ASN       string   `json:"asn"`
	ISP       string   `json:"isp"`
	Org       string   `json:"org"`
	Hostnames []string `json:"hostnames"`
	Domains   []string `json:"domains"`
	Port      int64    `json:"port"`
	Transport string   `json:"transport"`
	OS        string   `json:"os"`
	Data      string   `json:"data"`
	Product   string   `json:"product"`
	Version   string   `json:"version"`
	Info      string   `json:"info"`
	CPE       []string `json:"cpe"`
	CPE23     []string `json:"cpe23"`
	Tags      []string `json:"tags"`

	Shodan   Shodan   `json:"_shodan"`
	Location Location `json:"location"`
	Opts     Opts     `json:"opts"`
	Cloud    *Cloud   `json:"cloud"`
	DNS      *DNS     `json:"dns"`
	ISAKMP   *ISAKMP  `json:"isakmp"`
	NTP      *NTP     `json:"ntp"`
	HTTP     *HTTP    `json:"http"`
	SSH      *SSH     `json:"ssh"`
	SSL      *SSL     `json:"ssl"`
}

type Shodan struct {
	ID      string      `json:"id"`
	Crawler string      `json:"crawler"`
	Module  string      `json:"module"`
	Options interface{} `json:"options"`
	PTR     bool        `json:"ptr"`
}

type Opts struct {
	Raw        string        `json:"raw"`
	Heartbleed string        `json:"heartbleed"`
	Vulns      []interface{} `json:"vulns"`
}

type Location struct {
	AreaCode     string  `json:"area_code"`
	City         string  `json:"city"`
	CountryCode  string  `json:"country_code"`
	CountryCode3 string  `json:"country_code3"`
	CountryName  string  `json:"country_name"`
	DMACode      string  `json:"dma_code"`
	Latitude     float64 `json:"latitude"`
	Longitude    float64 `json:"longitude"`
	PostalCode   string  `json:"postal_code"`
	RegionCode   string  `json:"region_code"`
}

type DNS struct {
	Recursive        bool   `json:"recursive"`
	ResolverHostName string `json:"resolver_hostname"`
	ResolverID       string `json:"resolver_id"`
	Software         string `json:"software"`
}

type SSH struct {
	Cipher      string `json:"cipher"`
	Fingerprint string `json:"fingerprint"`
	Hash        string `json:"hash"`
	Kex         Kex    `json:"kex"`
	Key         string `json:"key"`
	MAC         string `json:"mac"`
	Type        string `json:"type"`
}

type Kex struct {
	CompressionAlgorithms   []string `json:"compression_algorithms"`
	EncryptionAlgorithms    []string `json:"encryption_algorithms"`
	KexFollows              bool     `json:"kex_follows"`
	KexAlgolithms           []string `json:"kex_algorithms"`
	Languages               []string `json:"languages"`
	MACAlgorithms           []string `json:"mac_algorithms"`
	ServerHostKeyAlgorithms []string `json:"server_host_key_algorithms"`
	Unused                  int64    `json:"unused"`
}

type HTTP struct {
	Components      map[string]interface{} `json:"components"` // e.g. jquery, bootstrap
	Location        string                 `json:"location"`
	Favicon         Favicon                `json:"favicon"`
	Host            string                 `json:"host"`
	HTML            string                 `json:"html"`
	HTMLHash        int64                  `json:"html_hash"`
	Redirects       []Redirect             `json:"redirects"`
	Robots          string                 `json:"robots"`
	RobotsHash      int64                  `json:"robots_hash"`
	SecurityTxt     string                 `json:"securitytxt"`
	SecurityTxtHash string                 `json:"securitytxt_hash"`
	Server          string                 `json:"server"`
	Sitemap         string                 `json:"sitemap"`
	SitemapHash     string                 `json:"sitemap_hash"`
	Title           string                 `json:"title"`
}

type Favicon struct {
	Hash     int64  `json:"hash"`
	Data     string `json:"data"`
	Location string `json:"location"`
}

type Redirect struct {
	Host     string `json:"host"`
	Data     string `json:"data"`
	Location string `json:"location"`
}

type SSL struct {
	Chain           []string     `json:"chain"`
	ChainSHA256     []string     `json:"chain_sha256"`
	Jarm            string       `json:"jarm"`
	Ja3s            string       `json:"ja3s"`
	DHParams        DHParams     `json:"dhparams"`
	Versions        []string     `json:"versions"`
	AcceptableCAS   []string     `json:"acceptable_cas"`
	TLSExt          []TLSExtData `json:"tlsext"`
	HandshakeStates []string     `json:"handshake_states"`
	ALPN            []string     `json:"alpn"`
	Cipher          Cipher       `json:"cipher"`
	Trust           Trust        `json:"trust"`
	OCSP            interface{}  `json:"ocsp"`
	Cert            Cert         `json:"cert"`
}

type DHParams struct {
	Prime     string `json:"prime"`
	PublicKey string `json:"public_key"`
	Bits      int64  `json:"bits"`
	Generator int64  `json:"generator"`
}

type TLSExtData struct {
	ID   int64  `json:"id"`
	Name string `json:"name"`
}

type Cipher struct {
	Version string `json:"version"`
	Bits    int64  `json:"bits"`
	Name    string `json:"name"`
}

type Trust struct {
	Revoked bool    `json:"revoked"`
	Browser Browser `json:"browser"`
}

type Browser struct {
	Apple     bool `json:"apple"`
	Mozilla   bool `json:"mozilla"`
	Microsoft bool `json:"microsoft"`
}

type Cert struct {
	SigAlg      string      `json:"sig_alg"`
	Issued      string      `json:"issued"`
	Expires     string      `json:"expires"`
	Expired     bool        `json:"expired"`
	Version     int64       `json:"version"`
	Extensions  []Extension `json:"extensions"`
	Fingerprint Fingerprint `json:"fingerprint"`
	Serial      BigInt      `json:"serial"`
	Subject     Subject     `json:"subject"`
	PubKey      PubKey      `json:"pubkey"`
	Issuer      Issuer      `json:"issuer"`
}

type Subject struct {
	CN string `json:"CN"`
}

type PubKey struct {
	Type string `json:"type"`
	Bits int64  `json:"bits"`
}

type Issuer struct {
	C  string `json:"C"`
	CN string `json:"CN"`
	O  string `json:"O"`
}

type Extension struct {
	Data string `json:"data"`
	Name string `json:"name"`
}

type Fingerprint struct {
	SHA256 string `json:"sha256"`
	SHA1   string `json:"sha1"`
}

type ISAKMP struct {
	InitiatorSPI string   `json:"initiator_spi"`
	ResponderSPI string   `json:"responder_spi"`
	MsgID        string   `json:"msg_id"`
	NextPayload  int64    `json:"next_payload"`
	ExchangeType int64    `json:"exchange_type"`
	Length       int64    `json:"length"`
	Version      string   `json:"version"`
	VendorIDs    []string `json:"vendor_ids"`
	Flags        Flags    `json:"flags"`
	Aggressive   *ISAKMP  `json:"aggressive"`
}

type Flags struct {
	Encryption     bool `json:"encryption"`
	Authentication bool `json:"authentication"`
	Commit         bool `json:"commit"`
}

type NTP struct {
	ClockOffset    float64     `json:"clock_offset"`
	Delay          float64     `json:"delay"`
	Monlist        interface{} `json:"monlist"`
	Leap           int64       `json:"leap"`
	Poll           int64       `json:"poll"`
	Precision      int64       `json:"precision"`
	Reftime        float64     `json:"reftime"`
	Refid          int64       `json:"refid"`
	RootDelay      float64     `json:"root_delay"`
	RootDispersion float64     `json:"root_dispersion"`
	Stratum        int64       `json:"stratum"`
	Version        int64       `json:"version"`
}

type Cloud struct {
	Region   string      `json:"region"`
	Service  interface{} `json:"service"`
	Provider string      `json:"provider"`
}

type BigInt struct {
	value string
}

func (i BigInt) String() string {
	return i.value
}

func (i *BigInt) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || string(b) == "null" {
		return nil
	}
	i.value = string(b)
	return nil
}
