# mbot

mbot is a command-line interface for [makerbotd](https://github.com/tjhorner/makerbotd). You can control your MakerBot 3D printers with it.

**This software is mid-development. It is not stable yet.** That said, feel free to play around with it!

## Usage

```
Usage: mbot <flags> <subcommand> <subcommand args>

Subcommands:
	cancel           Cancel the printer's current job.
	info             Get info about a printer.
	ls               List connected printers.
	pause            Suspend the printer's current job.
	print            Send a print file to a printer.
	resume           Resume the printer's current job.
	snapshot         Prints a snapshot of the printer's camera to stdout.
	status           Get the current job that a printer is running.
	timelapse        Record a time lapse of a printed object.
```

## Environment Variables

- `MBOT_PROTOCOL`: Either `unix` or `tcp` (default: `unix`)
- `MBOT_HOST`: Either the path to where the UNIX socket is listening or the base URL to where the daemon is listening (e.g. `http://localhost:6969`; default: `/var/run/makerbot.sock`)

## License

TBD