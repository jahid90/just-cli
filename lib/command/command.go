package command

import (
	"errors"
	"os"
	"os/exec"

	"github.com/jahid90/just/cmd/do/config/justfile"
)

// GeneratorFn A function to generate an exec.Cmd that can be run
type GeneratorFn func(string, *justfile.Just) (*exec.Cmd, error)

// Validate validates the command
func Validate(command string) error {

	// check that the command exists
	_, err := exec.LookPath(command)
	if err != nil {
		return errors.New("Error: " + command + " - command not found")
	}

	return nil
}

// Run Runs the command and attaches its stdout and stderr to os's stdout and stderr respectively
func Run(cmd *exec.Cmd, stdout *os.File, stderr *os.File) error {

	var err error

	// attach os stdout and stderr to cmd's stdout and stderr streams
	cmd.Stdout = stdout
	cmd.Stderr = stderr

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

	return nil
}
