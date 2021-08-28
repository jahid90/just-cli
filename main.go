package main

import (
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/fatih/color"
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
	app.Flags = []cli.Flag{}
	app.Commands = cmd.GetSubCommands()

	// Make doCmd the default when no subcommand is specified
	app.Action = do.Cmd().Action
	app.Flags = append(app.Flags, do.Cmd().Flags...)

	err := app.Run(os.Args)

	if err != nil {

		splits := strings.Split(err.Error(), ":")

		if len(splits) != 1 {
			t := splits[0]
			m := strings.Join(splits[1:], ":")

			// log.Fatal does not respect the color codes
			if strings.HasPrefix(err.Error(), "info") {
				info := color.New(color.FgGreen).SprintFunc()
				fmt.Println(info(t+":") + m)
			} else if strings.HasPrefix(err.Error(), "warn") {
				warn := color.New(color.FgYellow).SprintFunc()
				fmt.Println(warn(t+":") + m)
				os.Exit(1)
			} else {
				alert := color.New(color.FgRed).SprintFunc()
				fmt.Println(alert(t+":") + m)
				os.Exit(1)
			}
		} else {
			log.Fatal(err)
		}
	}
}
