package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"reflect"

	"github.com/google/subcommands"
)

type infoCmd struct {
	key string
}

func (*infoCmd) Name() string     { return "info" }
func (*infoCmd) Synopsis() string { return "Get info about a printer." }
func (*infoCmd) Usage() string {
	return `info [-key key] <printer>:
  Get info about a printer.
`
}

func (p *infoCmd) SetFlags(f *flag.FlagSet) {
	f.StringVar(&p.key, "key", "", "the info key to retrieve")
}

func (p *infoCmd) Execute(c context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ctx := c.(mbotCtx)

	pid := f.Args()[0]

	printer, err := ctx.Client.GetPrinter(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if p.key == "" {
		marshaled, err := json.MarshalIndent(*printer, "", "  ")
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		fmt.Println(string(marshaled))

		return subcommands.ExitSuccess
	}

	fmt.Println(reflect.ValueOf(*printer).FieldByName(p.key).String())
	return subcommands.ExitSuccess
}
