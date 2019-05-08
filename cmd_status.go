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

type statusCmd struct {
	raw bool
}

func (*statusCmd) Name() string     { return "status" }
func (*statusCmd) Synopsis() string { return "Get the current job that a printer is running." }
func (*statusCmd) Usage() string {
	return `status [--raw] <printer>:
  Get the current job that a printer is running.
`
}

func (p *statusCmd) SetFlags(f *flag.FlagSet) {
	f.BoolVar(&p.raw, "raw", false, "if true, will output raw JSON")
}

func (p *statusCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() == 0 {
		fmt.Println("Please provide the name or ID of the printer to get the status for.")
		return subcommands.ExitUsageError
	}

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

	if p.raw {
		marshaled, err := json.MarshalIndent(*job, "", "  ")
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		fmt.Println(string(marshaled))
	} else {
		tmpl, _ := template.New("currentjob").Parse(`Current Job:
  - Process Type: {{.Name}}
  - Progress: {{.Progress}}%
  - Step: {{.Step.Humanize}}
  - File Path: {{.Filepath}}
  - Cancellable: {{.Cancellable}}{{ if .Cancellable }}
    (hint: you can cancel with "mbot cancel"){{ end }}
`)

		tmpl.Execute(os.Stdout, *job)
	}

	return subcommands.ExitSuccess
}
