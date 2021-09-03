package executor

import (
	"os/exec"
	"time"

	"github.com/jahid90/just/core/logger"
)

// ExecuteMany Executes a slice of commands; stops when any one fails
func ExecuteMany(cmds []*exec.Cmd) error {

	if len(cmds) == 0 {
		logger.Warn("no commands passed")
		return nil
	}

	start := time.Now()

	for _, cmd := range cmds {
		err := execute(cmd)
		if err != nil {
			return err
		}
	}

	end := time.Now()
	logger.Infof("took %s", end.Sub(start).String())

	return nil
}

// Execute Runs a command
func Execute(cmd *exec.Cmd) error {

	start := time.Now()

	err := execute(cmd)
	if err != nil {
		return err
	}

	end := time.Now()
	logger.Infof("took %s", end.Sub(start).String())

	return nil
}

func execute(cmd *exec.Cmd) error {
	logger.Infof("executing %s", cmd.String())

	// start the command and await termination
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
