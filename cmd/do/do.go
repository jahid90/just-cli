package do

import (
	"fmt"
	"os/exec"

	"github.com/urfave/cli/v2"
)

func runCommand(cmd string, args ...string) ([]byte, error) {
	out, err := exec.Command(cmd, args...).Output()

	if err != nil {
		return nil, err
	}

	return out, nil
}

// Cmd A sub-command that prints a greeting
func Cmd() *cli.Command {
	return &cli.Command{
		Name:  "do",
		Usage: "Runs a command",
		Action: func(c *cli.Context) error {
			out, err := runCommand(c.Args().First(), c.Args().Tail()...)

			if err != nil {
				return err
			}

			fmt.Println(string(out))
			return nil
		},
	}
}
