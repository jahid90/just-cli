package main

import (
	"log"
	"os"

	"github.com/jahid90/just/cmd"
	"github.com/jahid90/just/cmd/do"
	"github.com/urfave/cli/v2"
)

func main() {

	// Disable timestamps in log messages
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	app := cli.NewApp()
	app.Name = "just"
	app.Usage = "A command runner"
	app.Description = "Runs commands defined by aliases in a config file.\n" +
		"Looks for a config file named just.json/just.yaml in the current directory.\n" +
		"A different config file can be provided using the `--config-file` switch\n" +
		"\n" +
		"Usage examples:\n" +
		"\tTo list the available commands, run `just --list`\n" +
		"\tTo execute a command, run `just <alias>`"

	app.Version = "1.0.0"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config-file",
			Aliases: []string{"c"},
			Usage:   "the config file to use",
		},
		&cli.StringFlag{
			Name:    "output",
			Aliases: []string{"o"},
			Usage:   "Print config as json/yaml",
		},
	}
	app.Commands = cmd.GetSubCommands()

	// Make doCmd the default when no subcommand is specified
	app.Action = do.Cmd().Action
	app.Flags = append(app.Flags, do.Cmd().Flags...)

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
