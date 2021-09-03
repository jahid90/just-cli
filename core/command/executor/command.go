package executor

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
)

// Validate validates the command
func Validate(command string) error {

	// check that the command exists
	_, err := exec.LookPath(command)
	if err != nil {
		return errors.New("error: " + command + " - command not found")
	}

	return nil
}

// Run Runs the command and attaches its stdout and stderr to os's stdout and stderr respectively
func Run(cmds []*exec.Cmd) error {

	for idx, cmd := range cmds {

		if idx != len(cmds)-1 {
			fmt.Println("just: dep@", cmd.String())
		} else {
			fmt.Println("just: main@", cmd.String())
		}

		var err error

		// attach os stdin, stdout and stderr to cmd's stdin, stdout and stderr streams
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		// start the command
		err = cmd.Start()
		if err != nil {
			return err
		}

		// wait till command's termination
		err = cmd.Wait()
		if err != nil {
			return err
		}
	}

	return nil
}
