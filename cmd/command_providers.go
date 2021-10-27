package main

import (
	"fmt"
	"sort"

	"github.com/mkideal/cli"
)

// parameters of 'providers' command.
type providersT struct {
	cli.Helper
}

func (a *providersT) Validate(ctx *cli.Context) error {
	return nil
}

var providersC = &cli.Command{
	Name: "providers",
	Desc: "Show supported provider types",
	Argv: func() interface{} { return new(providersT) },
	Fn:   execShowProviders,
}

func execShowProviders(ctx *cli.Context) error {
	argv := ctx.Argv().(*providersT)

	r := newShowProvidersRunner(*argv)
	return r.Run()
}

type ShowProvidersRunner struct{}

func newShowProvidersRunner(p providersT) ShowProvidersRunner {
	return ShowProvidersRunner{}
}

func (r *ShowProvidersRunner) Run() error {
	providers := make([]string, 0, len(providerMap))
	for key := range providerMap {
		providers = append(providers, key)
	}
	sort.Strings(providers)
	fmt.Printf("%v\n", providers)
	return nil
}
