package do

import (
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

			// parse the config
			config, err := parseConfig(c)
			if err != nil {
				return err
			}

			// handle flags
			if c.Bool("list") {
				err := config.GetListing()
				if err != nil {
					return err
				}

				return nil
			}

			// handle no args
			if c.Args().Len() == 0 {
				cli.ShowSubcommandHelp(c)
				return nil
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
