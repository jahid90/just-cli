package justfile

import (
	"errors"
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/jahid90/just/lib"
)

// Just A type representing a just config file
type Just struct {
	Version  string            `json:"version"`
	Commands map[string]string `json:"commands"`
}

// ShowListing Prints a table of the available commands
func (j *Just) ShowListing() error {

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
}

// LookupAlias Returns the command corrsponding to an alias
func (j *Just) LookupAlias(alias string) (string, error) {

	// check if the alias is present in the config file
	entry, ok := j.Commands[alias]
	if !ok {
		return "", errors.New("Error: alias `" + alias + "` not found in the config file")
	}

	return entry, nil
}

// ParserFn A function representing a parser
// On being invoked, parses the contents passed to it, generates a Just config file and returns a pointer to it
type ParserFn func([]byte) (*Just, error)

// GetParserFn Returns a parser function to parse the config file
func GetParserFn() ParserFn {

	return func(contents []byte) (*Just, error) {

		j := &Just{}
		err := lib.ParseJSON(contents, j)
		if err != nil {
			return nil, err
		}

		return j, nil
	}

}
