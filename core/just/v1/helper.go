package v1

import (
	"errors"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/command/executor"
	"github.com/jahid90/just/core/command/generator"
	"github.com/jahid90/just/core/just/api"
	"github.com/jahid90/just/core/logger"
)

var CommandGeneratorFn = func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error) {

	j, ok := config.(*Just)
	if !ok {
		logger.Error("bad type; expected v1 Just")
		return nil, errors.New("internal error")
	}

	aka, err := j.LookupAlias(alias)
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

	commandLine := strings.Split(entry, " ")
	cmd := commandLine[0]
	args := commandLine[1:]

	// generate the command; ignore any additional arguments supplied
	cmdExec := exec.Command(cmd, args...)

	return []*exec.Cmd{cmdExec}, nil
}

func GenerateApi(just *Just) *api.JustApi {
	return &api.JustApi{
		Version: func() string { return just.Version },
		Format:  just.Format,
		ShowListing: func(showShortListing bool) ([]byte, error) {
			if showShortListing {
				just.ShowShortListing()
				return nil, nil
			} else {
				just.ShowListing()
				return nil, nil
			}
		},
		Execute: func(alias string) error {
			for cmdAlias, cmd := range just.Commands {
				if cmdAlias == alias {

					var commandLine []string
					commandLine = append(commandLine, "-c")
					commandLine = append(commandLine, cmd)

					cmdExec := generator.Generate(nil, "sh", commandLine)

					e := executor.NewExecutor()
					e.AddExecutionUnit(cmdExec, "")
					if err := e.Execute(); err != nil {
						return err
					}

					return nil
				}
			}

			return errors.New("no alias matched")
		},
		ShowCommand: just.ShowCommand,
	}
}
