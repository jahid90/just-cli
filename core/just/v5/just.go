package v5

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/jahid90/just/core/misc"
	"gopkg.in/yaml.v2"
)

// Just A type representing a v5 just config
type Just struct {
	Version  string    `json:"version" yaml:"version"`
	Commands []Command `json:"commands" yaml:"commands"`
}

// Command A type representing a command
type Command struct {
	Alias       string   `json:"alias" yaml:"alias"`
	Exec        string   `json:"command" yaml:"command"`
	Description string   `json:"description" yaml:"description"`
	Depends     []string `json:"depends" yaml:"depends"`
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
	for _, cmd := range j.Commands {
		color.Yellow("> " + cmd.Alias)
		if len(cmd.Description) != 0 {
			fmt.Println(misc.Ellipsify("    "+cmd.Description, 80))
		}
		fmt.Println(misc.Ellipsify("    Execs: "+cmd.Exec, 80))
		if len(cmd.Depends) != 0 {
			fmt.Println("    Depends On: ")
			for _, dep := range cmd.Depends {
				fmt.Println("      - " + dep)
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
	fmt.Fprintln(w, "  "+strPrintBlue("ALIAS")+"\t\t"+strPrintBlue("COMMAND"))
	fmt.Fprintln(w, "  "+strPrintBlue("-----")+"\t\t"+strPrintBlue("-------"))
	for _, cmd := range j.Commands {
		fmt.Fprintln(w, "  "+strPrintYellow(cmd.Alias)+"\t\t"+strPrintWhite(misc.Ellipsify(cmd.Exec, 60)))
	}
	fmt.Fprintln(w)

	// flush the listing to output
	w.Flush()

	return nil
}

// LookupAlias Returns the command corresponding to an alias
func (j *Just) LookupAlias(alias string) (interface{}, error) {

	// check if the alias is present in the config file
	for _, cmd := range j.Commands {
		if cmd.Alias == alias {
			return cmd.Exec, nil
		}
	}

	return "", errors.New("error: alias `" + alias + "` not found in the config file")
}

// LookupDependencies Returns the dependent aliases of an alias
func (j *Just) LookupDependencies(alias string) ([]string, error) {

	for _, cmd := range j.Commands {
		if cmd.Alias == alias {
			return cmd.Depends, nil
		}
	}

	return []string{}, nil
}

// // Convert Converts config to v1-4
// func (j *Just) Convert() ([]byte, error) {

// 	v4 := &v1.Just{}
// 	v4.Version = "4"
// 	v4.Commands = make(map[string]string)

// 	for _, cmd := range j.Commands {
// 		v4.Commands[cmd.Alias] = cmd.Exec
// 	}

// 	return yaml.Marshal(v4)
// }

// ShowCommand Returns the command corresponding to an alias
func (j *Just) ShowCommand(alias string) (string, error) {

	for _, cmd := range j.Commands {
		if cmd.Alias == alias {
			return cmd.Exec, nil
		}
	}

	return "", errors.New("error: alias `" + alias + "` not found in the config file")

}
