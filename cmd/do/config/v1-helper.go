package config

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"

	"github.com/jahid90/just/lib"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/urfave/cli/v2"
)

func handleV1(contents []byte) (Config, error) {
	j, err := justfile.ParseV1(contents)
	if err != nil {
		return Config{}, nil
	}

	c, err := configFromV1(j)
	if err != nil {
		return Config{}, err
	}

	return c, nil
}

func configFromV1(j justfile.JustV1) (Config, error) {
	c := Config{
		RunCmd: func(c *cli.Context) error {
			cmd, err := getCmdV1(c, j)
			if err != nil {
				return err
			}

			err = lib.RunCommand(cmd)
			if err != nil {
				return err
			}

			return nil
		},
		GetListing: func() error {

			// handle no commands listed in the config file
			if len(j.Commands) == 0 {
				return errors.New("Error: no commands found in config file")
			}

			// format the listing in tabular form
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
			fmt.Fprintln(w, "Available commands are:")
			for alias, cmd := range j.Commands {
				fmt.Fprintln(w, "  "+alias+"\t"+cmd+"\t")
			}
			fmt.Fprintln(w)

			// flush the listing to output
			w.Flush()

			return nil
		},
	}

	return c, nil
}

func getCmdV1(c *cli.Context, j justfile.JustV1) (*exec.Cmd, error) {

	alias := c.Args().First()

	// check if the alias is present in the config file
	entry, ok := j.Commands[alias]
	if !ok {
		return nil, errors.New("Error: alias `" + alias + "` not found in the config file")
	}

	commandLine := strings.Split(entry, " ")
	command := commandLine[0]
	args := commandLine[1:]

	// check that the command exists
	_, err := exec.LookPath(command)
	if err != nil {
		return nil, errors.New("Error: " + command + " - command not found")
	}

	// output the command we are running
	fmt.Println("just @" + entry)

	// generate the command; ignore any additional arguments supplied
	cmd := exec.Command(command, args...)

	return cmd, nil
}
