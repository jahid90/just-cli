package cli

import (
	"log"
	"os"

	"github.com/jahid90/just/core/logger"
	"github.com/jahid90/just/input/cli/cmd"
	"github.com/jahid90/just/input/cli/cmd/do"
	"github.com/urfave/cli/v2"
)

func Run(gitVersion string) {

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
		"\tTo execute a command, run `just <alias>`" +
		"\n" +
		"\n" +
		"If no sub-command is passed, `do` is inferred."

	app.Version = "1.0.0 - " + gitVersion
	app.Flags = []cli.Flag{}
	app.Commands = cmd.GetSubCommands()

	// Make doCmd the default when no subcommand is specified
	app.Action = do.Cmd().Action
	app.Flags = append(app.Flags, do.Cmd().Flags...)

	err := app.Run(os.Args)

	if err != nil {
		logger.Fatal(err)
	}
}
