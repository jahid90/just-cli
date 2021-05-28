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
	app.Description = "Runs commands defined in a config file. Looks for a config file named just.json in the current directory. A different config file can be provided using the --config-file switch. \n\nSupports versions [1-4] for the config file. \n\tVer. 1 allows simple commands with no env vars or sub-command expansions. \n\tVer. 2 uses a stack based parsing login to identify and run sub-commands before running the main command. \n\tVer. 3 uses a lexer and a parser based on a grammar to parse the command and execute it. \n\tVer. 4 uses the underlying system's os shell to run the commands and supports both env vars and sub-command expansions. \nNote: Ver. 2 and Ver. 3 are not stable yet. \n\nUsage examples: \nTo list the available commands, run `just --list` \nTo execute a command, run `just <alias`"
	app.Version = "1.0.0"
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
