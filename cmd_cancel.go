package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/tjhorner/makerbotd/api"
)

type cancelCmd struct{}

func (*cancelCmd) Name() string     { return "cancel" }
func (*cancelCmd) Synopsis() string { return "Cancel the printer's current job." }
func (*cancelCmd) Usage() string {
	return `cancel <printer>:
  Cancel the printer's current job.
`
}

func (p *cancelCmd) SetFlags(f *flag.FlagSet) {}

func (p *cancelCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	client := args[0].(*api.Client)

	pid := f.Args()[0]

	_, err := client.CancelCurrentJob(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Current job stopped.")

	return subcommands.ExitSuccess
}
