package do

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"
	"text/tabwriter"

	"github.com/urfave/cli/v2"
)

func showListing(config *justV1) {

	// handle no commands listed in the config file
	if len(config.Commands) <= 0 {
		fmt.Println("No commands found in config file")
		return
	}

	// Format the listing in tabular form
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w, "Available commands are:")
	for alias, cmd := range config.Commands {
		fmt.Fprintln(w, "  "+alias+"\t"+cmd+"\t")
	}
	fmt.Fprintln(w)

	// flush the listing to output
	w.Flush()
}

func runCommand(cmd string, args ...string) ([]byte, error) {

	// execute the command and capture its stdout and stderr streams
	out, err := exec.Command(cmd, args...).CombinedOutput()
	if err != nil {
		return out, err
	}

	return out, nil
}

func handleCommand(config *justV1, arg string) error {

	// check if the command is present in the config file
	entry, ok := config.Commands[arg]
	if !ok {
		return errors.New("command `" + arg + "` not found in the config file")
	}

	// output the command we are running
	fmt.Println("just @" + entry)

	// execute the command; ignore any additional arguments supplied
	command := strings.Split(entry, " ")
	out, err := runCommand(command[0], command[1:]...)

	// print the output from the command run
	fmt.Println(string(out))

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
		Action: func(c *cli.Context) error {

			config, err := parseConfig()
			if err != nil {
				return err
			}

			arg := c.Args().First()
			switch arg {
			// TODO: use a flag
			case "list":
				showListing(&config)

				return nil
			default:
				err := handleCommand(&config, arg)
				if err != nil {
					return err
				}

				return nil
			}
		},
	}
}
