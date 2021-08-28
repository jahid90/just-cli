package justfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jahid90/just/lib"
	"gopkg.in/yaml.v2"
)

// JustV5 A type representing a v5 just config
type JustV5 struct {
	Version  string    `json:"version" yaml:"version"`
	Commands []Command `json:"commands" yaml:"commands"`
}

// Command A type representing a command in a JustV5 struct
type Command struct {
	Alias       string   `json:"alias" yaml:"alias"`
	Exec        string   `json:"command" yaml:"command"`
	Description string   `json:"description" yaml:"description"`
	Depends     []string `json:"depends" yaml:"depends"`
}

// Format Formats the config file into the requested format
func (j *JustV5) Format(format string) ([]byte, error) {
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
func (j *JustV5) ShowListing() error {
	if len(j.Commands) == 0 {
		return errors.New("warn: no commands found in config")
	}

	fmt.Println("Available commands are:")
	fmt.Println()
	for _, cmd := range j.Commands {
		fmt.Println("> " + cmd.Alias)
		if len(cmd.Description) != 0 {
			fmt.Println("    " + lib.Ellipsify(cmd.Description, 80))
		}
		fmt.Println("    Execs: " + lib.Ellipsify(cmd.Exec, 80))
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
func (j *JustV5) ShowShortListing() error {

	// handle no commands listed in the config file
	if len(j.Commands) == 0 {
		return errors.New("warn: no commands found in config")
	}

	// format the listing in tabular form
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Available commands are:")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "  ALIAS\t\tCOMMAND")
	fmt.Fprintln(w, "  -----\t\t-------")
	for _, cmd := range j.Commands {
		fmt.Fprintln(w, "  "+cmd.Alias+"\t\t"+lib.Ellipsify(cmd.Exec, 50))
	}
	fmt.Fprintln(w)

	// flush the listing to output
	w.Flush()

	return nil
}

// LookupAlias Returns the command corresponding to an alias
func (j *JustV5) LookupAlias(alias string) (string, error) {

	// check if the alias is present in the config file
	for _, cmd := range j.Commands {
		if cmd.Alias == alias {
			return cmd.Exec, nil
		}
	}

	return "", errors.New("error: alias `" + alias + "` not found in the config file")
}

// LookupDependencies Returns the dependent aliases of an alias
func (j *JustV5) LookupDependencies(alias string) ([]string, error) {

	for _, cmd := range j.Commands {
		if cmd.Alias == alias {
			return cmd.Depends, nil
		}
	}

	return []string{}, nil
}

// Convert Converts config to v4
func (j *JustV5) Convert() ([]byte, error) {

	v4 := &Just{}
	v4.Version = "4"
	v4.Commands = make(map[string]string)

	for _, cmd := range j.Commands {
		v4.Commands[cmd.Alias] = cmd.Exec
	}

	return yaml.Marshal(v4)
}
