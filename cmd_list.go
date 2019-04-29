package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type listCmd struct{}

func (*listCmd) Name() string     { return "ls" }
func (*listCmd) Synopsis() string { return "List connected printers." }
func (*listCmd) Usage() string {
	return `ls:
  List connected printers.
`
}

func (p *listCmd) SetFlags(f *flag.FlagSet) {}

func (p *listCmd) Execute(c context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ctx := c.(mbotCtx)

	printers, err := ctx.Client.GetPrinters()
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	for _, printer := range *printers {
		fmt.Printf("%s: %s (%s)\n", printer.Serial, printer.MachineName, printer.BotType)
	}

	return subcommands.ExitSuccess
}
