package justfile

import (
	"errors"
	"os/exec"
)

var CommandV6GeneratorFn = func(alias string, appendArgs []string, c *Config) ([]*exec.Cmd, error) {
	return nil, errors.New("warn: not yet implemented")
}
