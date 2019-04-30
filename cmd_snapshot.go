package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/google/subcommands"
	"github.com/tjhorner/makerbotd/api"
)

type snapshotCmd struct{}

func (*snapshotCmd) Name() string     { return "snapshot" }
func (*snapshotCmd) Synopsis() string { return "Prints a snapshot of the printer's camera to stdout." }
func (*snapshotCmd) Usage() string {
	return `snapshot <printer>:
  Prints a snapshot of the printer's camera to stdout.
`
}

func (p *snapshotCmd) SetFlags(f *flag.FlagSet) {}

func (p *snapshotCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Println("Please provide the name or ID of the printer to get a snapshot for.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Args()[0]

	snapshot, err := client.GetPrinterSnapshot(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	os.Stdout.Write(*snapshot)
	return subcommands.ExitSuccess
}
