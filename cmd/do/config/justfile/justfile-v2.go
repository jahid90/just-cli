package justfile

import (
	"github.com/jahid90/just/lib"
)

// JustV2 A type representing v2 of just config file
type JustV2 struct {
	Version  string    `json:"version"`
	Commands []Command `json:"commands"`
}

// Command Represents a command in v2
type Command struct {
	// The alias of the command that will be used to run it
	Alias string `json:"alias"`

	// The action to run when the alias is passed
	Action string `json:"action"`

	// Any env variables to pass onto the action
	Env map[string]string `json:"env"`
}

// ParseV2 parses JustV2
func ParseV2(contents []byte) (JustV2, error) {
	j := JustV2{}
	err := lib.ParseJSON(contents, &j)
	if err != nil {
		return JustV2{}, err
	}

	return j, nil
}
