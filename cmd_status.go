package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"

	"github.com/google/subcommands"
	"github.com/tjhorner/makerbotd/api"
)

type statusCmd struct{}

func (*statusCmd) Name() string     { return "status" }
func (*statusCmd) Synopsis() string { return "Get the current job that a printer is running." }
func (*statusCmd) Usage() string {
	return `status <printer>:
  Get the current job that a printer is running.
`
}

func (p *statusCmd) SetFlags(f *flag.FlagSet) {}

func (p *statusCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	client := args[0].(*api.Client)

	pid := f.Args()[0]

	job, err := client.GetCurrentJob(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if job.Name == "" {
		fmt.Println("There is currently no job running.")
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
