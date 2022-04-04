package v5

import (
	"errors"
	"os/exec"

	"github.com/jahid90/just/core/command/executor"
	"github.com/jahid90/just/core/command/generator"
	"github.com/jahid90/just/core/just/api"
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

func GenerateApi(config *Just) *api.JustApi {

	logger.Debug("generating the api")

	return &api.JustApi{
		Version: func() string {
			return config.Version
		},
		Format: config.Format,
		ShowListing: func(showShortListing bool) ([]byte, error) {
			if showShortListing {
				config.ShowShortListing()
			} else {
				config.ShowListing()
			}

			// must return instead of printing
			return nil, nil
		},
		Execute: func(alias string) error {
			command, err := findCommandMatching(alias, config)
			if err != nil {
				return err
			}
			logger.Debugf("alias matched a command: %#v", command)

			unit := generateExecFrom(command)
			e := executor.NewExecutor([]*executor.ExecutionUnit{unit})

			if err := e.Execute(); err != nil {
				return err
			}

			return nil
		},
		ShowCommand: config.ShowCommand,
	}
}

func findCommandMatching(alias string, config *Just) (*Command, error) {

	logger.Debug("finding command matching alias")

	for _, cmd := range config.Commands {
		if cmd.Alias == alias {
			return &cmd, nil
		}
	}

	return nil, errors.New("no such alias is defined in config")
}

func generateExecFrom(command *Command) *executor.ExecutionUnit {

	var commandLine []string
	commandLine = append(commandLine, "-c")
	commandLine = append(commandLine, command.Exec)

	cmd := generator.Generate(nil, "sh", commandLine)
	return executor.NewExecutionUnit(cmd, "")
}
