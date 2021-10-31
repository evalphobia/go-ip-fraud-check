go-ip-fraud-check
----

[![License: MIT][401]][402] [![GoDoc][101]][102] [![Release][103]][104] [![Build Status][201]][202] [![Coveralls Coverage][203]][204] [![Codecov Coverage][205]][206]
[![Go Report Card][301]][302] [![Code Climate][303]][304] [![BCH compliance][305]][306] [![CodeFactor][307]][308] [![codebeat][309]][310] [![Scrutinizer Code Quality][311]][312] [![FOSSA Status][403]][404]


<!-- Basic -->

[101]: https://godoc.org/github.com/evalphobia/go-ip-fraud-check?status.svg
[102]: https://godoc.org/github.com/evalphobia/go-ip-fraud-check
[103]: https://img.shields.io/github/release/evalphobia/go-ip-fraud-check.svg
[104]: https://github.com/evalphobia/go-ip-fraud-check/releases/latest
[105]: https://img.shields.io/github/downloads/evalphobia/go-ip-fraud-check/total.svg?maxAge=1800
[106]: https://github.com/evalphobia/go-ip-fraud-check/releases
[107]: https://img.shields.io/github/stars/evalphobia/go-ip-fraud-check.svg
[108]: https://github.com/evalphobia/go-ip-fraud-check/stargazers


<!-- Testing -->

[201]: https://github.com/evalphobia/go-ip-fraud-check/workflows/test/badge.svg
[202]: https://github.com/evalphobia/go-ip-fraud-check/actions?query=workflow%3Atest
[203]: https://coveralls.io/repos/evalphobia/go-ip-fraud-check/badge.svg?branch=master&service=github
[204]: https://coveralls.io/github/evalphobia/go-ip-fraud-check?branch=master
[205]: https://codecov.io/gh/evalphobia/go-ip-fraud-check/branch/master/graph/badge.svg
[206]: https://codecov.io/gh/evalphobia/go-ip-fraud-check


<!-- Code Quality -->

[301]: https://goreportcard.com/badge/github.com/evalphobia/go-ip-fraud-check
[302]: https://goreportcard.com/report/github.com/evalphobia/go-ip-fraud-check
[303]: https://codeclimate.com/github/evalphobia/go-ip-fraud-check/badges/gpa.svg
[304]: https://codeclimate.com/github/evalphobia/go-ip-fraud-check
[305]: https://bettercodehub.com/edge/badge/evalphobia/go-ip-fraud-check?branch=master
[306]: https://bettercodehub.com/
[307]: https://www.codefactor.io/repository/github/evalphobia/go-ip-fraud-check/badge
[308]: https://www.codefactor.io/repository/github/evalphobia/go-ip-fraud-check
[309]: https://codebeat.co/badges/142f5ca7-da37-474f-9264-f708ade08b5c
[310]: https://codebeat.co/projects/github-com-evalphobia-go-ip-fraud-check-master
[311]: https://scrutinizer-ci.com/g/evalphobia/go-ip-fraud-check/badges/quality-score.png?b=master
[312]: https://scrutinizer-ci.com/g/evalphobia/go-ip-fraud-check/?branch=master

<!-- License -->
[401]: https://img.shields.io/badge/License-MIT-blue.svg
[402]: LICENSE.md
[403]: https://app.fossa.com/api/projects/git%2Bgithub.com%2Fevalphobia%2Fgo-ip-fraud-check.svg?type=shield
[404]: https://app.fossa.com/projects/git%2Bgithub.com%2Fevalphobia%2Fgo-ip-fraud-check?ref=badge_shield


go-ip-fraud-check has a feature to detect fraud from ip addresss.
it provides both of cli binary and golang API.

# Supported Providers

- [AbuseIPDB](https://www.abuseipdb.com/)
- [Big Data Cloud](https://www.bigdatacloud.com/)
- [dbip](https://db-ip.com)
- [IP2Proxy](https://www.ip2location.com/web-service/ip2proxy)
- [ip-api.com](https://ip-api.com/)
- [ipdata](https://ipdata.co/)
- [ipgeolocation](https://ipgeolocation.io/)
- [ipify.org](https://www.ipify.org/)
- [ipinfo.io](https://ipinfo.io/)
- [IPQualityScore](https://www.ipqualityscore.com/)
- [Ipregistry](https://ipregistry.co/)
- [ipstack](https://ipstack.com/)
- [MaxMind minFraud](https://www.maxmind.com/en/solutions/minfraud-services/)
- [Shodan](https://shodan.io/)
- [Spur](https://spur.us/)


# Quick Usage for binary

## install

Download binary from release page, or build from source:

```bash
$ git clone --depth 1 https://github.com/evalphobia/go-ip-fraud-check.git
$ cd ./go-ip-fraud-check/cmd
$ go build -o ./go-ip-fraud-check .
```

## Subcommands

### root command

```bash
$ go-ip-fraud-check
Commands:

  help     show help
  single   Exec api call of ip address fraud check providers for single ip
  list     Exec api call of ip address fraud check providers from csv list file
  providers   Show supported provider types
```

### single command

`single` command is used to check single ip address.

```bash
./go-ip-fraud-check single -h

Exec api call of ip address fraud check providers for single ip

Options:

  -h, --help       display help information
  -p, --provider  *set types of api provider (space separated) --provider='ipdata ipinfo minfraud'
  -i, --ip         input ip address --ip='8.8.8.8'
      --route      set if you need route data from IRR --route
      --debug      set if you need verbose logs --debug
```

For example, you can check ip address like below

```bash
# set auth data
$ export FRAUD_CHECK_IPDATACO_APIKEY=xxx
$ export FRAUD_CHECK_IPINFOIO_TOKEN=yyy

# check ip address
$ ./go-ip-fraud-check single -p 'ipdata ipinfo' -i 8.8.8.8

2021/10/25 00:54:26 [INFO] Use ipdata.co
2021/10/25 00:54:26 [INFO] Use ipinfo.io
{"list":[{"service_name":"ipdata.co","ip":"8.8.8.8","hostname":"","isp":"Google LLC","organization":"","asn":15169,"risk_score":0,"is_vpn":false,"is_hosting":false,"is_proxy":false,"is_tor":false,"is_bot":false,"is_bogon":false,"has_other_threat":false,"threat_comment":"","country":"US","city":"","region":"","latitude":0,"longitude":0,"error":""},{"service_name":"ipinfo.io","ip":"8.8.8.8","hostname":"dns.google","isp":"","organization":"Google LLC","asn":15169,"risk_score":0,"is_vpn":false,"is_hosting":false,"is_proxy":false,"is_tor":false,"is_bot":false,"is_bogon":false,"has_other_threat":false,"threat_comment":"","country":"US","city":"Mountain View","region":"California","latitude":37.4056,"longitude":-122.0775,"error":""}],"as_prefix":null}
```

### list command

`list` command is used to check multiple ip address from list and save results to output file.

```bash
./go-ip-fraud-check list -h

Exec api call of ip address fraud check providers from csv list file

Options:

  -h, --help           display help information
  -p, --provider      *set types of api provider (space separated) --provider='ipdata ipinfo minfraud'
  -i, --input         *input csv/tsv file path --input='./input.csv'
  -o, --output        *output tsv file path --output='./output.tsv'
      --route          set if you need route data from IRR (this might be slow) --route
      --interval       time interval after a API call to handle rate limit (ms=msec s=sec, m=min) --interval=1.5s
  -m, --parallel[=2]   parallel number (multiple API calls) --parallel=2
  -v, --verbose        set if you need detail logs --verbose
      --debug          set if you use HTTP debug feature --debug
```

For example, you can check ip address from csv list like below

```bash
# set auth data
$ export FRAUD_CHECK_IPDATACO_APIKEY=xxx
$ export FRAUD_CHECK_IPINFOIO_TOKEN=yyy

# prepare CSV file
$ cat input.csv
ip_address
8.8.8.8
8.8.4.4
1.1.1.1


# check risk from the CSV file
$ ./go-ip-fraud-check list -p 'ipdata ipinfo' -i ./input.csv -o ./output.tsv
2021/10/25 00:58:29 [INFO] Use ipdata.co
2021/10/25 00:58:29 [INFO] Use ipinfo.io
2021/10/25 00:58:30 [INFO] exec #: [2]
2021/10/25 00:58:29 [INFO] exec #: [0]
2021/10/25 00:58:31 [INFO] exec #: [1]

$ cat output.tsv
service	ip_address	hostname	risk_score	isp	organization	asn	country	city	region	latitude	longitude	is_vpn	is_hosting	is_proxy	is_tor	is_bot	is_bogon	has_other_threat	threat_comment	error
ipdata.co	8.8.8.8		0.00000	Google LLC		15169	US			0.00000	0.00000	false	false	false	false	false	false	false
ipinfo.io	8.8.8.8	dns.google	0.00000		Google LLC	15169	US	Mountain View	California	37.40560	-122.07750	false	false	false	false	false	false	false
ipdata.co	8.8.4.4		0.00000	Google LLC		15169	US			0.00000	0.00000	false	false	false	false	false	false	false
ipinfo.io	8.8.4.4	dns.google	0.00000		Google LLC	15169	US	Mountain View	California	37.40560	-122.07750	false	false	false	false	false	false	false
ipdata.co	1.1.1.1		0.00000	Cloudflare, Inc.		13335	AU			0.00000	0.00000	false	false	false	false	false	false	false
ipinfo.io	1.1.1.1	one.one.one.one	0.00000		Cloudflare, Inc.	13335	US	San Francisco	California37.76210	-122.39710	false	true	false	false	false	false	false



# if provider has a rate limit, then use --interval and --parallel option.
$ ./go-ip-fraud-check list -p 'shodan' -i ./input.csv -o ./output.tsv --interval=1.2s --parallel=1
```


### providers command

`providers` command is used to see supported providers.

```bash
$ ./go-ip-fraud-check providers
[bigdatacloud ip2proxy ipdata ipgeolocation ipinfo ipqualityscore ipregistry minfraud shodan]
```

# Quick Usage for API

```go
package main

import (
	"fmt"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
	"github.com/evalphobia/go-ip-fraud-check/provider/ipdataco"
	"github.com/evalphobia/go-ip-fraud-check/provider/ipinfoio"
)

func main() {
	conf := ipfraudcheck.Config{
        // you can set auth values to config directly, otherwise used from environment variables.
		IPdatacoAPIKey:  "<your ipdata.co API key>",
		IPinfoioToken:  "<your ipinfo.io API token>",
		UseRoute:   true,
		Debug:      false,
	}

	svc, err := ipfraudcheck.New(conf, []provider.Provider{
		&ipdataco.IPdatacoProvider{},
		&ipinfo.IPinfoioProvider{},
	})
	if err != nil {
		panic(err)
	}

	// execute score API
	resp, err := svc.CheckIP("8.8.8.8")
	if err != nil {
		panic(err)
	}

	for _, r := range resp.List {
		// just print response in json format
		b, _ := json.Marshal(r)
		fmt.Printf("%s", string(b))
	}
}
```

see example dir for more examples.


# Environment variables

| Name | Description |
|:--|:--|
| `FRAUD_CHECK_ABUSEIPDB_APIKEY` | [AbuseIPDB API Key](https://docs.abuseipdb.com/). |
| `BIGDATACLOUD_APIKEY` | [Big Data Cloud API Key](https://www.bigdatacloud.com/sdk). |
| `FRAUD_CHECK_DBIP_APIKEY` | [dbip API key](https://db-ip.com/api/doc.php). |
| `FRAUD_CHECK_IP2PROXY_APIKEY` | [ip2proxy API key](https://www.ip2location.com/web-service/ip2proxy/). |
| `FRAUD_CHECK_IP2PROXY_PACKAGE` | [ip2proxy package parameter](https://www.ip2location.com/web-service/ip2proxy/). |
| `FRAUD_CHECK_IPDATACO_APIKEY` | [ipdata.co API key](https://docs.ipdata.co/). |
| `FRAUD_CHECK_IPGEOLOCATION_APIKEY` | [ipgeolocation API key](https://ipgeolocation.io/documentation.html). |
| `FRAUD_CHECK_IPIFY_APIKEY` | [ipify API key](https://geo.ipify.org/docs). |
| `FRAUD_CHECK_IPINFOIO_TOKEN` | [ipinfo.io API token](https://ipinfo.io/developers). |
| `IPQS_APIKEY` | [IPQualityScore API Key](https://www.ipqualityscore.com/documentation/overview). |
| `FRAUD_CHECK_IPSTACK_APIKEY` | [ipstack API key](https://ipstack.com/documentation). |
| `IPREGISTRY_APIKEY` | [Ipregistry API Key](https://ipregistry.co/docs/authentication). |
| `MINFRAUD_ACCOUNT_ID` | [MaxMind Account ID](https://support.maxmind.com/account-faq/license-keys/how-do-i-generate-a-license-key/). |
| `MINFRAUD_LICENSE_KEY` | [MaxMind License Key](https://support.maxmind.com/account-faq/license-keys/how-do-i-generate-a-license-key/). |
| `FRAUD_CHECK_SHODAN_APIKEY` | [Shodan API Key](https://developer.shodan.io/api/requirements). |
| `FRAUD_CHECK_SPUR_TOKEN` | [spur API token](https://spur.us/products/context-api). |
