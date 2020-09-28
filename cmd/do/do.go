package do

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "do",
		Usage: "Runs a command",
		Action: func(c *cli.Context) error {
			fmt.Println("Run the provided command")
			return nil
		},
	}
}
