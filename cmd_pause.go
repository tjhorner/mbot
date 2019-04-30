package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/tjhorner/makerbotd/api"
)

type pauseCmd struct{}

func (*pauseCmd) Name() string     { return "pause" }
func (*pauseCmd) Synopsis() string { return "Suspend the printer's current job." }
func (*pauseCmd) Usage() string {
	return `pause <printer>:
  Suspend the printer's current job.
`
}

func (p *pauseCmd) SetFlags(f *flag.FlagSet) {}

func (p *pauseCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Println("Please provide the name or ID of the printer to suspend the job for.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Arg(0)

	_, err := client.SuspendCurrentJob(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Current job stopped.")

	return subcommands.ExitSuccess
}
