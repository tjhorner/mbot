package main

import (
	"flag"
	"os"

	"github.com/tjhorner/makerbotd/api"

	"github.com/google/subcommands"
)

func main() {
	subcommands.Register(&jobCmd{}, "")
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

	ctx := mbotCtx{client}
	os.Exit(int(subcommands.Execute(ctx)))
}
