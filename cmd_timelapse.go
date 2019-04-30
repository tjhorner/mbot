package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"path/filepath"
	"time"

	"github.com/tjhorner/makerbot-rpc"

	"github.com/tjhorner/makerbotd/api"

	"github.com/google/subcommands"
)

type timelapseCmd struct {
	interval int
}

func (*timelapseCmd) Name() string     { return "timelapse" }
func (*timelapseCmd) Synopsis() string { return "Record a time lapse of a printed object." }
func (*timelapseCmd) Usage() string {
	return `timelapse [--interval seconds] <printer> <filepath>:
  Record a time lapse of a printed object.
`
}

func (p *timelapseCmd) SetFlags(f *flag.FlagSet) {
	f.IntVar(&p.interval, "interval", 5, "how often a snapshot should be taken, in seconds")
}

func (p *timelapseCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	client := args[0].(*api.Client)

	pid := f.Args()[0]
	path := f.Args()[1]

	fmt.Println("Sending print file...")

	_, err := client.Print(pid, path)
	if err != nil {
		fmt.Println(err)
		return subcommands.ExitFailure
	}

	fmt.Println("Waiting for print to start...")

	for {
		status, err := client.GetCurrentJob(pid)
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		if status.Step == makerbot.StepPrinting && status.Name == "PrintProcess" {
			break
		}

		fmt.Printf("process=%s step=%s progress=%d\n", status.Name, status.Step.String(), *status.Progress)

		time.Sleep(5 * time.Second)
	}

	td, err := ioutil.TempDir("", "mbot_timelapse")
	if err != nil {
		fmt.Println("An error occurred getting a temporary directory for the screenshots! Cancelling print.")
		fmt.Printf("%+v\n", err)
		client.CancelCurrentJob(pid)
		return subcommands.ExitFailure
	}

	fmt.Printf("Starting the time lapse (tmpdir=%s)...\n", td)

	for {
		status, err := client.GetCurrentJob(pid)
		if err != nil {
			fmt.Println(err)
			return subcommands.ExitFailure
		}

		fmt.Printf("process=%s step=%s progress=%d\n", status.Name, status.Step.String(), *status.Progress)

		ss, err := client.GetPrinterSnapshot(pid)
		if err != nil {
			fmt.Printf("Error getting snapshot: %+v\n", err)
		}

		fp := filepath.Join(td, fmt.Sprintf("%d.jpg", time.Now().UnixNano()))

		ioutil.WriteFile(fp, *ss, 0664)

		if status.Step == makerbot.StepCompleted {
			break
		}

		time.Sleep(time.Duration(p.interval) * time.Second)
	}

	fmt.Println("Time lapse completed! I didn't actually add the necessary code to compile all the screenshots, but you can do this to make an mp4 yourself if you have ffmpeg installed:")

	fmt.Printf("\n  cd %s && ffmpeg -r 5 -pattern_type glob -i '*.jpg' -c:v libx264 -r 30 -pix_fmt yuv420p out.mp4\n\n", td)

	fmt.Println("Then you'll have out.mp4 in that directory. Ta-da.")

	return subcommands.ExitSuccess
}
