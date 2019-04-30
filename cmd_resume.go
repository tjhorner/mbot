package main

import (
	"context"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/tjhorner/makerbotd/api"
)

type resumeCmd struct{}

func (*resumeCmd) Name() string     { return "resume" }
func (*resumeCmd) Synopsis() string { return "Resume the printer's current job." }
func (*resumeCmd) Usage() string {
	return `resume <printer>:
  Resume the printer's current job.
`
}

func (p *resumeCmd) SetFlags(f *flag.FlagSet) {}

func (p *resumeCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Println("Please provide the name or ID of the printer to resume the job for.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Arg(0)

	_, err := client.ResumeCurrentJob(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Current job stopped.")

	return subcommands.ExitSuccess
}
