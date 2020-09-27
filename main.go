package main

import (
	"fmt"
	"log"
	"os"

	"github.com/jahid90/just/cmd/hello"
	"github.com/urfave/cli/v2"
)

func main() {
	app := &cli.App{
		Name:  "just",
		Usage: "A command runner that runs commands defined in a config file (Justfile by default)",
		Action: func(c *cli.Context) error {
			fmt.Println("Show usage here...")
			return nil
		},
	}

	app.Commands = hello.GetCommands()

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}
