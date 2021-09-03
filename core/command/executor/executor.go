package executor

import (
	"os/exec"

	"github.com/jahid90/just/core/logger"
)

// ExecuteMany Executes a slice of commands; stops when any one fails
func ExecuteMany(cmds []*exec.Cmd) error {

	for _, cmd := range cmds {
		err := Execute(cmd)
		if err != nil {
			return err
		}
	}

	return nil
}

// Execute Runs a command
func Execute(cmd *exec.Cmd) error {
	logger.Infof("executing %s", cmd.String())

	// start the command and await termination
	err := cmd.Run()
	if err != nil {
		return err
	}

	return nil
}
