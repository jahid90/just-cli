package v6

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/command/executor"
	"github.com/jahid90/just/core/command/generator"
	"github.com/jahid90/just/core/just/api"
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

func GenerateApi(config *Just) *api.JustApi {
	return &api.JustApi{
		Version: func() string { return config.Version },
		Format:  config.Format,
		ShowListing: func(showShortListing bool) ([]byte, error) {
			if showShortListing {
				config.ShowShortListing()
				return nil, nil
			} else {
				config.ShowListing()
				return nil, nil
			}
		},
		Execute: func(alias string) error {
			command, err := findCommandMatching(alias, config)
			if err != nil {
				return err
			}
			logger.Debugf("alias matched a command: %#v", command)

			cmds := generateExecStepsFrom(command)
			executor.ExecuteMany(cmds)

			return nil
		},
	}
}

func findCommandMatching(alias string, config *Just) (*Command, error) {
	for cmdAlias, cmd := range config.Commands {
		if alias == cmdAlias {
			return &cmd, nil
		}
	}

	return nil, errors.New("no alias matched")
}

func generateExecStepsFrom(command *Command) []*exec.Cmd {
	cmds := []*exec.Cmd{}

	for _, step := range command.Steps {
		if len(step.Uses) != 0 {
			logger.Info("skipping step as it uses action")
			continue
		}

		split := strings.Split(step.Run, " ")

		cmd := generator.Generate(step.Env, split[0], split[1:])
		cmds = append(cmds, cmd)
	}

	return cmds
}
