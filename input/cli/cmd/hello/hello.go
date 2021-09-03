package hello

import (
	"os"

	"github.com/jahid90/just/core/logger"
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
				logger.Info("Hello Stranger!")
			} else {
				logger.Infof("Hello, %s!", user)
			}

			return nil
		},
	}
}
