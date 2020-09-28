package hello

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "hello",
		Usage: "Says hello",
		Action: func(c *cli.Context) error {
			user, ok := os.LookupEnv("USER")

			if !ok {
				fmt.Println("Hello Stranger!")
			} else {
				fmt.Printf("Hello, %s!\n", user)
			}

			return nil
		},
	}
}
