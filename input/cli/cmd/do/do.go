package do

import (
	"fmt"

	"github.com/jahid90/just/core/just"
	"github.com/jahid90/just/core/just/api"
	"github.com/urfave/cli/v2"
)

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "do",
		Usage: "Runs a command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "config-file",
				Aliases: []string{"c"},
				Usage:   "the config file to use",
			},
			&cli.BoolFlag{
				Name:    "list",
				Aliases: []string{"l"},
				Usage:   "list the available commands",
			},
			&cli.BoolFlag{
				Name:    "short",
				Aliases: []string{"s"},
				Usage:   "list a short version of the available commands",
			},
			&cli.StringFlag{
				Name:    "output",
				Aliases: []string{"o"},
				Usage:   "Print config as json/yaml",
			},
			&cli.BoolFlag{
				Name:  "convert",
				Usage: "Convert config files between different versions",
			},
		},
		Action: func(c *cli.Context) error {

			// no args and no flags set; this means just the base command was run; show help and exit
			if c.Args().Len() == 0 && c.NumFlags() == 0 {
				cli.ShowAppHelp(c)

				return nil
			}

			// get the config file name
			configFile, err := getConfigFileName(c)
			if err != nil {
				return err
			}

			// get the api client
			api, err := just.GetApi(configFile)
			if err != nil {
				return err
			}

			// handle flags
			err = handleFlags(c, api)
			if err != nil {
				return err
			}

			// return if there are no args to run
			if c.Args().Len() == 0 {
				return nil
			}

			// run the command corresponding to the alias
			err = api.Execute(c.Args().First())
			if err != nil {
				return err
			}

			return nil
		},
	}
}

func handleFlags(c *cli.Context, api *api.JustApi) error {

	// handle list flag
	if c.Bool("list") {

		// doesn't return anything yet
		_, err := api.ShowListing(false)
		if err != nil {
			return err
		}

		return nil
	}

	// handle list short flag
	if c.Bool("short") {
		_, err := api.ShowListing(true)
		if err != nil {
			return err
		}

		return nil
	}

	// handle output flag
	if len(c.String("output")) != 0 {

		formatted, err := api.Format(c.String("output"))
		if err != nil {
			return err
		}

		fmt.Println(string(formatted))

		return nil
	}

	// if c.Bool("convert") {
	// 	converted, err := api.Convert()
	// 	if err != nil {
	// 		return err
	// 	}

	// 	fmt.Println(string(converted))

	// 	return nil
	// }

	return nil
}
