package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
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

func (p *cancelCmd) Execute(c context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ctx := c.(mbotCtx)

	pid := f.Args()[0]

	_, err := ctx.Client.CancelCurrentJob(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Current job stopped.")

	return subcommands.ExitSuccess
}
