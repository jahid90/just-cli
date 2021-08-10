package justfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"gopkg.in/yaml.v2"
)

// JustV5 A type representing a v5 just config
type JustV5 struct {
	Version  string    `json:"version" yaml:"version"`
	Commands []Command `json:"commands" yaml:"commands"`
}

// Command A type representing a command in a JustV5 struct
type Command struct {
	Alias       string `json:"alias" yaml:"alias"`
	Exec        string `json:"command" yaml:"command"`
	Description string `json:"description" yaml:"description"`
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

// ShowListing Prints a table of the available commands
func (j *JustV5) ShowListing() error {

	// handle no commands listed in the config file
	if len(j.Commands) == 0 {
		return errors.New("warn: no commands found in config file")
	}

	// format the listing in tabular form
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Available commands are:")
	fmt.Fprintln(w)
	fmt.Fprintln(w, "  ALIAS\t\tCOMMAND")
	fmt.Fprintln(w, "  -----\t\t-------")
	for _, cmd := range j.Commands {
		fmt.Fprintln(w, "  "+cmd.Alias+"\t\t"+cmd.Exec+"\t")
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
