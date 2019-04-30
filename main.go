package main

import (
	"context"
	"flag"
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

	sockPath, ok := os.LookupEnv("MBOT_SOCKET_PATH")
	if !ok {
		sockPath = api.DefaultUnixSocketPath
	}

	client := api.NewClientSocket(sockPath)

	ctx := context.Background()
	os.Exit(int(subcommands.Execute(ctx, client)))
}
