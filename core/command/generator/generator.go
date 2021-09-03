package generator

import (
	"os"
	"os/exec"
)

func Generate(env []string, command string, args []string) *exec.Cmd {
	cmd := exec.Command(command, args...)
	cmd.Env = env
	cmd.Stdin = os.Stdin
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	return cmd
}
