package command

import (
	"os"
	"os/exec"

	"github.com/jahid90/just/cmd/do/config/justfile"
)

// GeneratorFn A function to generate an exec.Cmd that can be run
type GeneratorFn func(string, *justfile.Just) (*exec.Cmd, error)

// RunCommand Runs the command and attaches its stdout and stderr to os's stdout and stderr respectively
func RunCommand(cmd *exec.Cmd) error {

	var err error

	// attach os stdout and stderr to cmd's stdout and stderr streams
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

	return nil
}
