package main

import (
	"fmt"
	"strings"

	"github.com/evalphobia/go-ip-fraud-check/provider"
	"github.com/evalphobia/go-ip-fraud-check/provider/ipdataco"
	"github.com/evalphobia/go-ip-fraud-check/provider/ipinfoio"
	"github.com/evalphobia/go-ip-fraud-check/provider/minfraud"
)

const (
	providerIPdataco = "ipdata"
	providerIPinfoio = "ipinfo"
	providerMinFraud = "minfraud"
)

var providerMap = map[string]provider.Provider{
	providerIPdataco: &ipdataco.IPdatacoProvider{},
	providerIPinfoio: &ipinfoio.IPinfoioProvider{},
	providerMinFraud: &minfraud.MinFraudProvider{},
}

func validateProviderString(s string) error {
	keys := make([]string, 0, len(providerMap))
	for k := range providerMap {
		keys = append(keys, k)
	}

	for _, v := range getProviderListString(s) {
		if _, ok := providerMap[v]; !ok {
			return fmt.Errorf("provider should be one of the %v, but [%s] is used", keys, v)
		}
	}
	return nil
}

func getProviderListString(s string) []string {
	var providers []string
	for _, v := range strings.Split(s, " ") {
		v = strings.TrimSpace(v)
		if v == "" {
			continue
		}
		providers = append(providers, v)
	}
	return providers
}

func getProvidersFromString(s string) ([]provider.Provider, error) {
	strList := getProviderListString(s)
	list := make([]provider.Provider, len(strList))
	for i, v := range strList {
		p, err := getProvider(v)
		if err != nil {
			return nil, err
		}
		list[i] = p
	}
	return list, nil
}

func getProvider(s string) (provider.Provider, error) {
	if p, ok := providerMap[s]; ok {
		return p, nil
	}
	return nil, fmt.Errorf("unknown provider for [%s]", s)
}