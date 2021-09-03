package justfile

import (
	"errors"
	"fmt"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/command/executor"
	v6 "github.com/jahid90/just/core/justfile/v6"
)

var CommandV6GeneratorFn = func(alias string, appendArgs []string, c *Config) ([]*exec.Cmd, error) {

	cmds := []*exec.Cmd{}

	// first run any dependent commands;
	deps, _ := c.LookupDependencies(alias)
	for _, dep := range deps {
		fmt.Println("depends on: " + dep)
	}

	aka, err := c.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	entry, ok := aka.(v6.Command)
	if !ok {
		return nil, errors.New("error: internal - unexpected type received")
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		fmt.Println("warn: ignoring extra arguments provided - " + strings.Join(appendArgs, ", "))
	}

	for _, step := range entry.Steps {

		if len(step.Uses) > 0 {
			fmt.Println("info: execute action: " + step.Uses)
			continue
		}

		program := strings.Split(step.Run, " ")[0]
		if err := executor.Validate(program); err != nil {
			fmt.Println("warn: command not found - " + program)
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
