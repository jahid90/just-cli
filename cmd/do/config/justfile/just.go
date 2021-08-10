package justfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"gopkg.in/yaml.v2"
)

// Just A type representing a just config
type Just struct {
	Version  string            `json:"version" yaml:"version"`
	Commands map[string]string `json:"commands" yaml:"commands"`
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

// ShowListing Prints a table of the available commands
func (j *Just) ShowListing() error {

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
	for alias, cmd := range j.Commands {
		fmt.Fprintln(w, "  "+alias+"\t\t"+cmd+"\t")
	}
	fmt.Fprintln(w)

	// flush the listing to output
	w.Flush()

	return nil
}

// LookupAlias Returns the command corresponding to an alias
func (j *Just) LookupAlias(alias string) (string, error) {

	// check if the alias is present in the config file
	entry, ok := j.Commands[alias]
	if !ok {
		return "", errors.New("error: alias `" + alias + "` not found in the config file")
	}

	return entry, nil
}
