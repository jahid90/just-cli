package v6

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/jahid90/just/core/misc"
	"github.com/jahid90/just/output/console/colorize"
	"gopkg.in/yaml.v2"
)

// Just A type representing a v6 just config
type Just struct {
	Version  string             `json:"version"`
	Commands map[string]Command `json:"commands"`
}

// Command A type representing a command
// A command can depend on other commands via the 'needs' directive
// A command can have multiple steps
type Command struct {
	Description string   `json:"description"`
	Needs       []string `json:"needs"`
	Steps       []Step   `json:"steps"`
}

// Step A type representing each step of a command
// A step can either use a pre-defined step action via the 'uses' directive
// or define the step using a name, env vars and the command to run
type Step struct {
	Uses string   `json:"uses"`
	Name string   `json:"name"`
	Env  []string `json:"env"`
	Run  string   `json:"run"`
}

// Format Formats the config file into the requested format
func (j *Just) Format(format string) ([]byte, error) {
	if format == "json" {
		formatted, err := json.MarshalIndent(j, "", "  ")
		if err != nil {
			return nil, err
		}

		return formatted, nil
	}

	if format == "yaml" {
		formatted, err := yaml.Marshal(j)
		if err != nil {
			return nil, err
		}

		return formatted, nil
	}

	return nil, errors.New("error: output must be one of ['json', 'yaml']")
}

// ShowListing Prints a list of the avaliable commands
func (j *Just) ShowListing() error {
	if len(j.Commands) == 0 {
		return errors.New("warn: no commands found in config")
	}

	fmt.Println()
	fmt.Println("Available commands are:")
	fmt.Println()
	for alias, cmd := range j.Commands {
		colorize.Print(*color.New(color.FgYellow), "> "+alias)
		if len(cmd.Description) != 0 {
			fmt.Println(misc.Ellipsify("  "+cmd.Description, 80))
		}
		if len(cmd.Needs) != 0 {
			fmt.Println()
			fmt.Println("  Depends On:")
			for _, dep := range cmd.Needs {
				color.Yellow("    - " + dep)
			}
		}
		if len(cmd.Steps) != 0 {
			fmt.Println()
			fmt.Println("  Steps:")
			for _, step := range cmd.Steps {
				if len(step.Uses) != 0 {
					fmt.Println("    - " + step.Uses)
				} else {
					fmt.Println("    - " + step.Name)
					if len(step.Env) != 0 {
						fmt.Println("      Env:")
						for _, env := range step.Env {
							fmt.Println("        - " + env)
						}
					}
					fmt.Println("      Run: " + step.Run)
				}
			}
		}
		fmt.Println()
	}

	return nil
}

// ShowShortListing Prints a table of the available commands
func (j *Just) ShowShortListing() error {

	// we have to wrap every print in tabwriter as color basically
	// adds color codes to the string to print them on the screen
	// and tabwriter will incorrectly assume the color ctrings to be
	// of different length with the inclusion of the codes
	strPrintWhite := color.New(color.FgWhite).SprintFunc()
	strPrintBlue := color.New(color.FgBlue).SprintFunc()
	strPrintYellow := color.New(color.FgYellow).SprintFunc()

	// handle no commands listed in the config file
	if len(j.Commands) == 0 {
		return errors.New("warn: no commands found in config")
	}

	// format the listing in tabular form
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w)
	fmt.Fprintln(w, strPrintWhite("Available commands are:"))
	fmt.Fprintln(w)
	// two blues to match coloring of alias
	fmt.Fprintln(w, "  "+strPrintBlue("ALIAS")+"\t\t"+strPrintBlue("DESCRIPTION"))
	fmt.Fprintln(w, "  "+strPrintBlue("-----")+"\t\t"+strPrintBlue("-----------"))
	for alias, cmd := range j.Commands {
		fmt.Fprintln(w, "  "+strPrintYellow(alias)+"\t\t"+strPrintWhite(misc.Ellipsify(cmd.Description, 60)))
	}
	fmt.Fprintln(w)

	// flush the listing to output
	w.Flush()

	return nil
}

// LookupAlias Returns the command corresponding to an alias
func (j *Just) LookupAlias(alias string) (interface{}, error) {

	// check if the alias is present in the config file
	for aka, cmd := range j.Commands {
		if aka == alias {
			// TODO - fix this
			return cmd, nil
		}
	}

	return "", errors.New("error: alias `" + alias + "` not found in the config file")
}

// LookupDependencies Returns the dependent aliases of an alias
func (j *Just) LookupDependencies(alias string) ([]string, error) {

	for aka, cmd := range j.Commands {
		if aka == alias {
			return cmd.Needs, nil
		}
	}

	return []string{}, nil
}

// Convert Converts config to v4
func (j *Just) Convert() ([]byte, error) {
	return nil, errors.New("warn: not supported")
}

// ShowCommand Returns the command(s) corresponding to an alias
func (j *Just) ShowCommand(alias string) (string, error) {

	for aka, cmd := range j.Commands {
		if aka == alias {

			var command []string
			for _, step := range cmd.Steps {
				command = append(command, step.Run)
			}

			return strings.Join(command, "\n"), nil
		}
	}

	return "", errors.New("error: alias `" + alias + "` not found in the config file")

}
