package hello

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "hello",
		Usage: "Says hello",
		Action: func(c *cli.Context) error {
			fmt.Println("Hello there!")
			return nil
		},
	}
}
