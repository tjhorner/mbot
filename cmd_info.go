package main

import (
	"context"
	"encoding/json"
	"flag"
	"fmt"
	"html/template"
	"os"

	"github.com/google/subcommands"
	"github.com/tjhorner/makerbotd/api"
)

type infoCmd struct {
	raw bool
}

func (*infoCmd) Name() string     { return "info" }
func (*infoCmd) Synopsis() string { return "Get info about a printer." }
func (*infoCmd) Usage() string {
	return `info [--raw] <printer>:
  Get info about a printer.
`
}

func (p *infoCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.raw, "raw", false, "if true, will output raw JSON")
}

func (p *infoCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Println("Please provide the name or ID of the printer to get info for.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Args()[0]

	printer, err := client.GetPrinter(pid)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	if p.raw {
		marshaled, err := json.MarshalIndent(*printer, "", "  ")
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		fmt.Println(string(marshaled))
	} else {
		tmpl, _ := template.New("printerinfo").Parse(`Printer Info: {{.MachineName}}
  - Bot Type: {{.BotType}}
  - Serial Number: {{.Serial}}
  - IP: {{.IP}}
  - JSON-RPC Port: {{.Port}}
  - Firmware Version: {{.FirmwareVersion.Major}}.{{.FirmwareVersion.Minor}}.{{.FirmwareVersion.Bugfix}} ({{.FirmwareVersion.Build}})
`)

		tmpl.Execute(os.Stdout, *printer)
	}

	return subcommands.ExitSuccess
}
