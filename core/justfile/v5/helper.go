package v5

import (
	"errors"
	"os/exec"

	"github.com/jahid90/just/core/logger"
)

var CommandGeneratorFn = func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error) {

	cmds := []*exec.Cmd{}

	j, ok := config.(*Just)
	if !ok {
		logger.Error("bad type; expect v5 Just")
		return nil, errors.New("internal error")
	}

	// first run any dependent commands;
	// TODO: handle transitive dependencies and cycles
	deps, _ := j.LookupDependencies(alias)
	for _, dep := range deps {
		aa, err := j.LookupAlias(dep)
		if err != nil {
			return nil, err
		}

		a, ok := aa.(string)
		if !ok {
			logger.Error("bad type, expected string")
			return nil, errors.New("internal error")
		}

		cmd := exec.Command("sh", "-c", a)
		cmds = append(cmds, cmd)
	}

	aka, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	entry, ok := aka.(string)
	if !ok {
		logger.Error("bad type, expected string")
		return nil, errors.New("internal error")
	}

	cmd := exec.Command("sh", "-c", entry)
	cmds = append(cmds, cmd)

	return cmds, nil
}
