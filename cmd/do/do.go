package do

import (
	"fmt"

	"github.com/jahid90/just/cmd/do/config"
	"github.com/urfave/cli/v2"
)

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "do",
		Usage: "Runs a command",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list the available commands",
			},
		},
		Action: func(c *cli.Context) error {

			// no args and no flags set; this means just the base command was run; show help and exit
			if c.Args().Len() == 0 && c.NumFlags() == 0 {
				cli.ShowAppHelp(c)

				return nil
			}

			// parse the config
			config, err := parseConfig(c)
			if err != nil {
				return err
			}

			// handle flags
			err = handleFlags(c, config)
			if err != nil {
				return err
			}

			// run the command
			err = config.RunCmd(c)
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func handleFlags(c *cli.Context, config *config.Config) error {

	// handle list flag
	if c.Bool("list") {
		err := config.GetListing()
		if err != nil {
			return err
		}

		return nil
	}

	// handle output flag
	if len(c.String("output")) != 0 {

		formatted, err := config.Format(c.String("output"))
		if err != nil {
			return err
		}

		fmt.Println(string(formatted))

		return nil
	}

	return nil
}
