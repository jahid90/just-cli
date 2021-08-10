package justfile

import "errors"

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

func (j *JustV5) ShowListing() error {
	return errors.New("warn: not yet implemented")
}

func (j *JustV5) LookupAlias(alias string) (string, error) {
	return "", errors.New("warn: not yet implemented")
}
