package justfile

import (
	"errors"
	"os/exec"
	"strings"
)

var CommandV4GeneratorFn = func(alias string, appendArgs []string, c *Config) ([]*exec.Cmd, error) {

	aka, err := c.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	entry, ok := aka.(string)
	if !ok {
		return nil, errors.New("error: internal - unexpected type received")
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	cmd := exec.Command("sh", "-c", entry)

	return []*exec.Cmd{cmd}, nil
}
