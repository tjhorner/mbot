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
func (*snapshotCmd) Synopsis() string { return "Takes a snapshot of the printer's camera." }
func (*snapshotCmd) Usage() string {
	return `snapshot <printer> <outfile>:
  Takes a snapshot of the printer's camera.
`
}

func (p *snapshotCmd) SetFlags(f *flag.FlagSet) {}

func (p *snapshotCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() <= 1 {
		fmt.Println("Please provide the name or ID of the printer to get a snapshot for as well as where to write the file.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Arg(0)

	snapshot, err := client.GetPrinterSnapshot(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fn := f.Arg(1)

	var stream *os.File
	if fn == "-" {
		stream = os.Stdout
	} else {
		stream, err = os.Create(fn)
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}
	}

	stream.Write(*snapshot)
	return subcommands.ExitSuccess
}
