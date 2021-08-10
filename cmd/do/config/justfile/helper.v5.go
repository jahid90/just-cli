package justfile

import (
	"os/exec"
	"strings"
)

var CommandV5GeneratorFn = func(alias string, appendArgs []string, c *Config) ([]*exec.Cmd, error) {

	cmds := []*exec.Cmd{}

	// first run any dependent commands;
	// TODO: handle transitive dependencies and cycles
	deps, _ := c.LookupDependencies(alias)
	for _, dep := range deps {
		a, err := c.LookupAlias(dep)
		if err != nil {
			return nil, err
		}

		cmd := exec.Command("sh", "-c", a)
		cmds = append(cmds, cmd)
	}

	entry, err := c.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	cmd := exec.Command("sh", "-c", entry)
	cmds = append(cmds, cmd)

	return cmds, nil
}
