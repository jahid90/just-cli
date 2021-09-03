package v2

import (
	"errors"
	"fmt"
	"io"
	"os/exec"
	"strings"

	"github.com/jahid90/just/core/just/api"
	v1 "github.com/jahid90/just/core/just/v1"
	"github.com/jahid90/just/core/logger"
	"github.com/jahid90/just/core/misc"
)

var CommandGeneratorFn = func(alias string, appendArgs []string, config interface{}) ([]*exec.Cmd, error) {

	j, ok := config.(*v1.Just)
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

	s := misc.NewRuneStack()
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

func GenerateApi(config *v1.Just) *api.JustApi {
	return v1.GenerateApi(config)
}
