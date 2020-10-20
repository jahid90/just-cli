package config

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"strings"

	"github.com/jahid90/just/lib/lexer"

	"github.com/jahid90/just/cmd/do/config/justfile"
	"github.com/jahid90/just/lib/command"
)

var commandV2GeneratorFn = func(alias string, j *justfile.Just) (*exec.Cmd, error) {

	entry, err := j.LookupAlias(alias)
	if err != nil {
		return nil, err
	}

	// output the command we are running
	fmt.Println("just @" + entry)

	c, args, env, err := parseCommandLine(entry)
	if err != nil {
		return nil, err
	}

	err = command.Validate(c)
	if err != nil {
		return nil, err
	}

	// generate the command; ignore any additional arguments supplied
	cmd := exec.Command(c, args...)
	cmd.Env = append(os.Environ(), env...)

	return cmd, nil
}

// parseCommandLine Parses a command line and generates (command, []arg, []env, error)
func parseCommandLine(commandLine string) (string, []string, []string, error) {

	var command string
	var args, env []string

	s := lexer.NewTokenStack()

	l := lexer.NewLexer(strings.NewReader(commandLine))
	buffer := l.Run()
	buffer.Print()

	// Pass0 _ Processes expressions
	err := pass0(buffer, s)
	if err != nil {
		return "", nil, nil, err
	}
	fmt.Println("Pass0 complete")

	// Pass1 - Parses env variables
	s.Reverse()
	env, err = pass1(s)
	if err != nil {
		return "", nil, nil, err
	}
	fmt.Println("Pass1 complete")
	fmt.Println(env)

	// Pass2 - Parses command
	command, err = pass2(s)
	if err != nil {
		return "", nil, nil, err
	}
	fmt.Println("Pass2 complete")
	fmt.Println(command)

	// Pass3 - Parses args
	args, err = pass3(s)
	if err != nil {
		return "", nil, nil, err
	}
	fmt.Println("Pass3 complete")
	fmt.Println(args)

	return command, args, env, nil
}

func pass0(buffer *lexer.TokenBuffer, s *lexer.TokenStack) error {

	for {

		if !buffer.HasNext() {
			break
		}

		token := buffer.Next()

		if token.IsExprStart() {

			// get the entire expression
			var cl []string
			for {
				ok := buffer.HasNext()
				if !ok {
					return errors.New("Error: stream consumed before expression could be completely parsed at pos: " + fmt.Sprint(token.Position))
				}

				t := buffer.Next()
				if t.IsExprEnd() {
					break
				}

				cl = append(cl, t.Value)

			}

			// exec and get result
			newCl := strings.Join(cl, " ")
			// fmt.Println("Found an expression: " + newCl)
			c, a, e, err := parseCommandLine(newCl)
			if err != nil {
				return err
			}

			cmd := exec.Command(c, a...)
			cmd.Env = append(os.Environ(), e...)

			// cmd.Run()
			out, err := cmd.Output()
			if err != nil {
				return err
			}

			token = &lexer.Token{Type: lexer.IDENT, Position: token.Position, Value: string(out)}

		}

		s.Push(token)
	}

	return nil
}

func pass1(s *lexer.TokenStack) ([]string, error) {

	// fmt.Println("== Pass1 ==")
	// s.Print()

	var env []string

	newStack := lexer.NewTokenStack()

	for {

		empty := s.IsEmpty()
		if empty {
			break
		}

		token, err := s.Pop()
		if err != nil {
			return nil, err
		}

		if !s.IsEmpty() {
			nextToken, err := s.Top()
			if err != nil {
				return nil, err
			}

			if nextToken.IsAssign() {
				// env variable found
				e, err := processEnv(token, s)
				if err != nil {
					return nil, err
				}

				// fmt.Println("Found an env: " + e)
				// s.Print()

				env = append(env, e)

				continue
			}
		}

		if token.IsComma() {
			// just ignore the comma, we'll find the next env var anyway
			continue
		}

		newStack.Push(token)
	}

	// newStack.Print()
	newStack.Reverse()
	*s = *newStack

	return env, nil
}

func processEnv(prev *lexer.Token, s *lexer.TokenStack) (string, error) {

	// fmt.Println("== Process Env ==")
	// fmt.Println("prev: " + prev.Value)

	// consume the '='
	_, err := s.Pop()
	if err != nil {
		return "", err
	}

	nextToken, err := s.Pop()
	if err != nil {
		return "", err
	}

	//fmt.Println("next: " + nextToken.Value)

	return prev.Value + "=" + nextToken.Value, nil
}

func pass2(s *lexer.TokenStack) (string, error) {

	fmt.Println("== Pass2 ==")
	s.Print()

	token, err := s.Pop()
	if err != nil {
		return "", errors.New("Error: stream ended before a command was found")
	}

	return token.Value, nil
}

func pass3(s *lexer.TokenStack) ([]string, error) {

	var args []string

	for {
		empty := s.IsEmpty()
		if empty {
			break
		}

		a, err := processArg(s)
		if err != nil {
			return nil, err
		}
		args = append(args, a)
	}

	return args, nil
}

func processArg(s *lexer.TokenStack) (string, error) {

	token, err := s.Pop()
	if err != nil {
		return "", err
	}

	return token.Value, nil
}
