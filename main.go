package main

import (
	"log"
	"os"

	"github.com/jahid90/just/cmd"
	"github.com/urfave/cli/v2"
)

func main() {

	// Disable timestamps in log messages
	log.SetFlags(log.Flags() &^ (log.Ldate | log.Ltime))

	app := &cli.App{
		Name:  "just",
		Usage: "A command runner that runs commands defined in a config file (just.json by default)",
	}

	app.Commands = cmd.GetSubCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
