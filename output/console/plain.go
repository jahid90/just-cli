package console

import (
	"fmt"
	"io"
)

type plainConsole struct {
	out io.Writer
}

func NewPlainConsole(out io.Writer) plainConsole {
	return plainConsole{
		out: out,
	}
}

func (p plainConsole) Print(args ...interface{}) error {
	_, err := fmt.Fprint(p.out, args...)

	if err != nil {
		return err
	}

	return nil
}

func (p plainConsole) Println(args ...interface{}) error {
	_, err := fmt.Fprintln(p.out, args...)

	if err != nil {
		return err
	}

	return nil
}

func (p plainConsole) Printf(format string, args ...interface{}) error {
	_, err := fmt.Fprintf(p.out, format, args...)

	if err != nil {
		return err
	}

	return nil
}

func (p plainConsole) Sprint(args ...interface{}) string {
	res := fmt.Sprint(args...)

	return res
}

func (p plainConsole) Sprintln(args ...interface{}) string {
	res := fmt.Sprintln(args...)

	return res
}

func (p plainConsole) Sprintf(format string, args ...interface{}) string {
	res := fmt.Sprintf(format, args...)

	return res
}
