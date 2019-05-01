package main

import (
	"context"
	"flag"
	"fmt"
	"os"

	"github.com/tjhorner/makerbotd/api"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(subcommands.HelpCommand(), "")
	subcommands.Register(subcommands.FlagsCommand(), "")
	subcommands.Register(subcommands.CommandsCommand(), "")

	subcommands.Register(&resumeCmd{}, "")
	subcommands.Register(&pauseCmd{}, "")
	subcommands.Register(&timelapseCmd{}, "")
	subcommands.Register(&statusCmd{}, "")
	subcommands.Register(&listCmd{}, "")
	subcommands.Register(&snapshotCmd{}, "")
	subcommands.Register(&infoCmd{}, "")
	subcommands.Register(&cancelCmd{}, "")
	subcommands.Register(&printCmd{}, "")
	flag.Parse()

	connProto, ok := os.LookupEnv("MBOT_PROTOCOL")
	if !ok {
		connProto = "unix"
	}

	connHost, ok := os.LookupEnv("MBOT_HOST")
	if !ok {
		connHost = api.DefaultUnixSocketPath
	}

	var client *api.Client
	if connProto == "tcp" {
		client = api.NewClientTCP(connHost)
	} else if connProto == "unix" {
		client = api.NewClientSocket(connHost)
	} else {
		panic(fmt.Errorf("invalid MBOT_PROTOCOL (got: %s, wanted: unix or tcp)", connProto))
	}

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx, client)))
}
