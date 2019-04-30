package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/tjhorner/makerbotd/api"

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

func (p *printCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() <= 1 {
		fmt.Println("Please provide the name or ID of the printer to print to as well as the path to the .makerbot file.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Args()[0]
	path := f.Args()[1]

	fmt.Println("Sending print file...")

	_, err := client.Print(pid, path)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Now printing.")

	return subcommands.ExitSuccess
}
