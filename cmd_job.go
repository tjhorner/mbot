package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/google/subcommands"
)

type jobCmd struct{}

func (*jobCmd) Name() string     { return "job" }
func (*jobCmd) Synopsis() string { return "Get the current job that a printer is doing." }
func (*jobCmd) Usage() string {
	return `job <printer>:
  Get the current job that a printer is doing.
`
}

func (p *jobCmd) SetFlags(f *flag.FlagSet) {}

func (p *jobCmd) Execute(c context.Context, f *flag.FlagSet, _ ...interface{}) subcommands.ExitStatus {
	ctx := c.(mbotCtx)

	pid := f.Args()[0]

	job, err := ctx.Client.GetCurrentJob(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if job.Name == "" {
		fmt.Print("There is currently no job running.")
		return subcommands.ExitSuccess
	}

	marshaled, err := json.MarshalIndent(*job, "", "  ")
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println(string(marshaled))

	return subcommands.ExitSuccess
}
