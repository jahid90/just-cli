package do

import (
	"os"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func runCommand(cmd *exec.Cmd) error {

	var err error

	// attach os stdout and stderr to cmd's stdout and stderr streams
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	// start the command
	err = cmd.Start()
	if err != nil {
		return err
	}

	// wait till command's termination
	err = cmd.Wait()
	if err != nil {
		return err
	}

	return nil
}

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

			// handle actual command
			cmd, err := config.GetCmd(c)
			if err != nil {
				return err
			}

			runCommand(cmd)

			return nil
		},
	}
}
