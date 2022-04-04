package executor

import (
	"os/exec"

	"github.com/jahid90/just/core/logger"
)

type ExecutionUnit struct {
	cmd         *exec.Cmd
	description string
}

func NewExecutionUnit(cmd *exec.Cmd, description string) *ExecutionUnit {
	return &ExecutionUnit{
		cmd:         cmd,
		description: description,
	}
}

func (e *ExecutionUnit) execute() error {
	if e.description != "" {
		logger.Info(e.description)
	}

	logger.Infof("executing %s", e.cmd.String())

	// start the command and await termination
	err := e.cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
