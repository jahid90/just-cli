package justfile

import (
	"errors"
	"os/exec"
)

var CommandV5GeneratorFn = func(alias string, appendArgs []string, c *Config) ([]*exec.Cmd, error) {

	cmds := []*exec.Cmd{}

	// first run any dependent commands;
	// TODO: handle transitive dependencies and cycles
	deps, _ := c.LookupDependencies(alias)
	for _, dep := range deps {
		aa, err := c.LookupAlias(dep)
		if err != nil {
			return nil, err
		}

		a, ok := aa.(string)
		if !ok {
			return nil, errors.New("error: internal - unexpected type received")
		}

		cmd := exec.Command("sh", "-c", a)
		cmds = append(cmds, cmd)
	}

	aka, err := c.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	entry, ok := aka.(string)
	if !ok {
		return nil, errors.New("error: internal - unexpected type received")
	}

	cmd := exec.Command("sh", "-c", entry)
	cmds = append(cmds, cmd)

	return cmds, nil
}
