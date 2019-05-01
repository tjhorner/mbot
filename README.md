# mbot

mbot is a command-line interface for [makerbotd](https://github.com/tjhorner/makerbotd). You can control your MakerBot 3D printers with it.

**This software is mid-development. It is not stable yet.** That said, feel free to play around with it!

## Usage

If your makerbotd listens on a different path, you can provide it with `MBOT_SOCKET_PATH`. TCP HTTP connections not currently supported.

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

## License

TBD