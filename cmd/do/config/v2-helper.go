package config

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/urfave/cli/v2"
)

func handleV2(contents []byte) (Config, error) {
	j, err := justfile.ParseV2(contents)
	if err != nil {
		return Config{}, nil
	}

	c, err := configFromV2(j)
	if err != nil {
		return Config{}, nil
	}

	return c, nil
}

func configFromV2(j justfile.JustV2) (Config, error) {
	c := Config{
		GetCmd: func(c *cli.Context) (*exec.Cmd, error) {

			alias := c.Args().First()

			// check if the alias is present in the config file
			var ok = false
			var entry justfile.Command
			for _, command := range j.Commands {
				if command.Alias == alias {
					ok = true
					entry = command
					break
				}
			}
			if !ok {
				return nil, errors.New("Error: alias `" + alias + "` not found in the config file")
			}

			commandLine := strings.Split(entry.Action, " ")
			command := commandLine[0]
			args := commandLine[1:]

			// check that the command exists
			_, err := exec.LookPath(command)
			if err != nil {
				return nil, errors.New("Error: " + command + " - command not found")
			}

			// output the command we are running
			fmt.Println("just @" + getEnv(entry.Env) + entry.Action)

			// generate the command; ignore any additional arguments supplied
			cmd := exec.Command(command, args...)
			cmd.Env = append(os.Environ(), getEnv(entry.Env))

			return cmd, nil
		},
		GetListing: func() error {

			// handle no commands listed in config file
			if len(j.Commands) == 0 {
				return errors.New("Error: no commands found in config file")
			}

			// format the listing in tabular form
			w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
			fmt.Fprintln(w, "Available commands are:")
			for _, command := range j.Commands {
				env := getEnv(command.Env)
				fmt.Fprintln(w, "  "+command.Alias+"\t"+env+command.Action+"\t")
			}
			fmt.Fprintln(w)

			// flush the listing to output
			w.Flush()

			return nil
		},
	}

	return c, nil
}

func getEnv(environment map[string]string) string {
	var env string
	for e, v := range environment {
		env += e + "=" + v
	}
	if len(env) > 0 {
		env += " "
	}

	return env
}
