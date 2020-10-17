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

func runCommand(cmd string, args ...string) error {

	var err error

	// execute the command and attach os stdout and stderr to its stdout and stderr streams
	command := exec.Command(cmd, args...)
	command.Stdout = os.Stdout
	command.Stderr = os.Stderr

	err = command.Start()
	if err != nil {
		return err
	}

	err = command.Wait()
	if err != nil {
		return err
	}

	return nil
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
	err := runCommand(command[0], command[1:]...)
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
