package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type printCmd struct{}

func (*printCmd) Name() string     { return "print" }
func (*printCmd) Synopsis() string { return "Send a print file to a printer." }
func (*printCmd) Usage() string {
	return `print <printer> <filepath>:
  Send a print file to a printer.
`
}

func (p *printCmd) SetFlags(f *flag.FlagSet) {}

func (p *printCmd) Execute(c context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ctx := c.(mbotCtx)

	pid := f.Args()[0]
	path := f.Args()[1]

	fmt.Println("Sending print file...")

	_, err := ctx.Client.Print(pid, path)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Now printing.")

	return subcommands.ExitSuccess
}
