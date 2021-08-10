package justfile

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/jahid90/just/lib"
)

var CommandV2GeneratorFn = func(alias string, appendArgs []string, j *Config) ([]*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// add any additional arguments provided
	if len(appendArgs) > 0 {
		entry = entry + " " + strings.Join(appendArgs, " ")
	}

	reduced, err := parseCommandLine(entry)
	if err != nil {
		return nil, err
	}
	fmt.Println(reduced)

	// c, e, a, err := split(reduced)
	// if err != nil {
	// 	return nil, err
	// }

	// err = command.Validate(c)
	// if err != nil {
	// 	return nil, err
	// }

	// generate the command; ignore any additional arguments supplied
	cmd := exec.Command("sh", "-c", reduced)

	return []*exec.Cmd{cmd}, nil
}

func parseCommandLine(input string) (string, error) {

	s := lib.NewRuneStack()
	reader := strings.NewReader(input)

	for {
		r, _, err := reader.ReadRune()
		if err != nil {
			if err == io.EOF {
				return s.AsString(), nil
			}

			return "", err
		}

		if r == ')' {
			// found an expression; evaluate it
			expr := ""
			for {
				i, err := s.Top()
				if err != nil {
					return "", err
				}

				if i != '(' {
					s.Pop()
					expr = string(i) + expr
				} else {
					break
				}
			}

			i, err := s.Pop()
			if err != nil {
				return "", err
			}
			if i != '(' {
				return "", errors.New("error: could not find start of expression: (")
			}

			i, err = s.Pop()
			if err != nil {
				return "", err
			}
			if i != '$' {
				return "", errors.New("error: could not find start of expression: $")
			}

			cmdOutput, err := exec.Command("sh", "-c", expr).Output()
			if err != nil {
				return "", err
			}
			in := strings.NewReader(string(cmdOutput))
			for {
				rr, _, err := in.ReadRune()
				if err != nil {
					if err == io.EOF {
						break
					}

					return "", err
				}
				s.Push(rr)
			}

			continue

		}

		s.Push(r)
	}
}

func split(commandline string) (string, []string, []string, error) {
	return "", nil, nil, errors.New("warn: not yet implemented")
}
