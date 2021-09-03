package v6

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/logger"
)

var CommandGeneratorFn = func(alias string, appendArgs []string, jj interface{}) ([]*exec.Cmd, error) {

	cmds := []*exec.Cmd{}

	j, ok := jj.(*Just)
	if !ok {
		logger.Error("bad type; expected v6 Just")
		return nil, errors.New("internal error")
	}

	// first run any dependent commands;
	deps, _ := j.LookupDependencies(alias)
	for _, dep := range deps {
		logger.Info("depends on: " + dep)
	}

	aka, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	entry, ok := aka.(Command)
	if !ok {
		logger.Error("bad type; expected string")
		return nil, errors.New("internal error")
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		logger.Warn("ignoring extra arguments provided - " + strings.Join(appendArgs, ", "))
	}

	for _, step := range entry.Steps {

		if len(step.Uses) > 0 {
			logger.Info("execute action: " + step.Uses)
			continue
		}

		var list []string = []string{}
		list = append(list, "-c")
		if len(step.Env) > 0 {
			list = append(list, strings.Join(step.Env, ","))
		}
		list = append(list, step.Run)
		cmd := exec.Command("sh", list...)
		cmds = append(cmds, cmd)
	}

	return cmds, nil
}
