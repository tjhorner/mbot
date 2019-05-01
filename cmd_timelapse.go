package main

import (
	"context"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"strings"
	"time"

	"github.com/tjhorner/makerbot-rpc"

	"github.com/tjhorner/makerbotd/api"

	"github.com/google/subcommands"
)

type timelapseCmd struct {
	interval int
	fps      int
	outFile  string
}

func (*timelapseCmd) Name() string     { return "timelapse" }
func (*timelapseCmd) Synopsis() string { return "Record a time lapse of a printed object." }
func (*timelapseCmd) Usage() string {
	return `timelapse [--interval seconds] [--fps fps] [-o outfile] <printer> <filepath>:
  Record a time lapse of a printed object.
`
}

func (p *timelapseCmd) SetFlags(f *flag.FlagSet) {
	var wd string
	wd, err := os.Getwd()
	if err != nil {
		wd = "/tmp"
	}

	f.IntVar(&p.interval, "interval", 5, "how often a snapshot should be taken, in seconds")
	f.IntVar(&p.fps, "fps", 20, "the fps of the final timelapse")
	f.StringVar(&p.outFile, "o", fmt.Sprintf("%s/timelapse_%d.mp4", wd, time.Now().Unix()), "output video file")
}

func (p *timelapseCmd) Execute(c context.Context, f *flag.FlagSet, args ...interface{}) subcommands.ExitStatus {
	if f.NArg() <= 1 {
		fmt.Println("Please provide the name or ID of the printer to print to as well as the path to the .makerbot file.")
		return subcommands.ExitUsageError
	}

	client := args[0].(*api.Client)

	pid := f.Arg(0)
	path := f.Arg(1)

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

	// invoke ffmpeg (I am too lazy to use libav)
	fmt.Println("Print is done! Compiling snapshots into mp4 file using ffmpeg...")

	ffmpegArgs := []string{"-r", strconv.Itoa(p.fps), "-pattern_type", "glob", "-i", "*.jpg", "-c:v", "libx264", "-r", strconv.Itoa(p.fps), "-pix_fmt", "yuv420p", p.outFile}
	cmd := exec.Command("ffmpeg", ffmpegArgs...)

	cmd.Dir = td

	err = cmd.Run()
	if err != nil {
		fullCmd := fmt.Sprintf("ffmpeg %s", strings.Join(ffmpegArgs, " "))
		fmt.Printf("Couldn't convert to mp4. Do you have ffmpeg installed? If you do, run this command inside of %s:\n\n    %s", td, fullCmd)
		fmt.Println("\nIf you don't have ffmpeg installed, you can grab it here: https://ffmpeg.org")
		return subcommands.ExitSuccess
	}

	fmt.Printf("Time lapse completed! Saved to: %s\n", p.outFile)

	err = os.RemoveAll(td)
	if err != nil {
		fmt.Printf("warn: could not remove tmp directory; you can try deleting it yourself: %s", td)
	}

	return subcommands.ExitSuccess
}
