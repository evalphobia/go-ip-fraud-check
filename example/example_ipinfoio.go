package main

import (
	"flag"
	"fmt"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/provider"
	"github.com/evalphobia/go-ip-fraud-check/provider/ipinfoio"
)

// nolint
func main() {
	var ipaddr string
	flag.StringVar(&ipaddr, "ipaddr", "", "set target ip address")
	flag.Parse()

	svc, err := ipfraudcheck.New(ipfraudcheck.Config{
		UseRoute: true,
	}, []provider.Provider{
		&ipinfoio.IPinfoioProvider{},
	})
	if err != nil {
		panic(err)
	}

	resp, err := svc.CheckIP(ipaddr)
	if err != nil {
		panic(err)
	}
	fmt.Printf("[%+v]\n", resp)
}
