package hello

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

/*
GetCommands returns the subcommands for the app.
*/
func GetCommands() []*cli.Command {

	helloCmd := hello()

	var subCommands []*cli.Command

	return append(subCommands, helloCmd)
}

func hello() *cli.Command {
	return &cli.Command{
		Name:  "hello",
		Usage: "Says hello",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello there!")
			return nil
		},
	}
}
