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
	app.Usage = "A command runner that runs commands defined in a config file (just.json by default)"
	app.Flags = []cli.Flag{
		&cli.StringFlag{
			Name:    "config-file",
			Aliases: []string{"c"},
			Usage:   "the config file to use",
			Value:   "just.json",
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
