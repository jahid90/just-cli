package justfile

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/fatih/color"
	"github.com/jahid90/just/lib"
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

	// we have to wrap every print in tabwriter as color basically
	// adds color codes to the string to print them on the screen
	// and tabwriter will incorrectly assume the color ctrings to be
	// of different length with the inclusion of the codes
	strPrintWhite := color.New(color.FgWhite).SprintFunc()
	strPrintBlue := color.New(color.FgBlue).SprintFunc()
	strPrintYellow := color.New(color.FgYellow).SprintFunc()

	// handle no commands listed in the config file
	if len(j.Commands) == 0 {
		return errors.New("warn: no commands found in config file")
	}

	// format the listing in tabular form
	w := tabwriter.NewWriter(os.Stdout, 0, 0, 3, ' ', tabwriter.TabIndent)
	fmt.Fprintln(w)
	fmt.Fprintln(w, strPrintWhite("Available commands are:"))
	fmt.Fprintln(w)
	fmt.Fprintln(w, "  "+strPrintBlue("ALIAS")+"\t\t"+strPrintBlue("COMMAND"))
	fmt.Fprintln(w, "  "+strPrintBlue("-----")+"\t\t"+strPrintBlue("-------"))
	for alias, cmd := range j.Commands {
		fmt.Fprintln(w, "  "+strPrintYellow(alias)+"\t\t"+strPrintWhite(lib.Ellipsify(cmd, 60)))
	}
	fmt.Fprintln(w)

	// flush the listing to output
	w.Flush()

	return nil
}

func (j *Just) ShowShortListing() error {
	return j.ShowListing()
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

// LookupDependencies Returns the dependent aliases of an alias
func (j *Just) LookupDependencies(alias string) ([]string, error) {
	return nil, errors.New("error: not supported")
}

// Convert Converts config to v5
func (j *Just) Convert() ([]byte, error) {

	v5 := &JustV5{}
	v5.Version = "5"
	v5.Commands = []CommandV5{}

	for alias, exec := range j.Commands {
		v5.Commands = append(v5.Commands, CommandV5{Alias: alias, Exec: exec, Description: "", Depends: []string{}})
	}

	return yaml.Marshal(v5)
}
