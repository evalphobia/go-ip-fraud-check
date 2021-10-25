package main

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/mkideal/cli"

	"github.com/evalphobia/go-ip-fraud-check/ipfraudcheck"
	"github.com/evalphobia/go-ip-fraud-check/log"
)

// parameters of 'single' command.
type singleT struct {
	cli.Helper
	Provider  string `cli:"*p,provider" usage:"set types of api provider (space separated) --provider='ipdata ipinfo minfraud'"`
	IPAddress string `cli:"i,ip" usage:"input ip address --ip='8.8.8.8'"`
	UseRoute  bool   `cli:"route" usage:"set if you need route data from IRR --route"`
	Debug     bool   `cli:"debug" usage:"set if you need verbose logs --debug"`
}

func (a *singleT) Validate(ctx *cli.Context) error {
	if a.IPAddress == "" {
		return errors.New("you must set --ip")
	}

	return validateProviderString(a.Provider)
}

var singleC = &cli.Command{
	Name: "single",
	Desc: "Exec api call of ip address fraud check providers for single ip",
	Argv: func() interface{} { return new(singleT) },
	Fn:   execSingle,
}

func execSingle(ctx *cli.Context) error {
	argv := ctx.Argv().(*singleT)

	r := newSingleRunner(*argv)
	return r.Run()
}

type SingleRunner struct {
	// parameters
	Provider  string
	IPAddress string
	UseRoute  bool
	Debug     bool
}

func newSingleRunner(p singleT) SingleRunner {
	return SingleRunner{
		Provider:  p.Provider,
		IPAddress: p.IPAddress,
		UseRoute:  p.UseRoute,
		Debug:     p.Debug,
	}
}

func (r *SingleRunner) Run() error {
	providerList, err := getProvidersFromString(r.Provider)
	if err != nil {
		panic(err)
	}

	svc, err := ipfraudcheck.New(ipfraudcheck.Config{
		UseRoute: r.UseRoute,
		Debug:    r.Debug,
		Logger:   &log.StdLogger{},
	}, providerList)
	if err != nil {
		panic(err)
	}

	resp, err := svc.CheckIP(r.IPAddress)
	if err != nil {
		panic(err)
	}

	b, err := json.Marshal(resp)
	if err != nil {
		return err
	}

	// just print response in json format
	fmt.Printf("%s\n", string(b))
	return nil
}
