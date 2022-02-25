package console

import (
	"io"

	"github.com/fatih/color"
)

type coloredConsole struct {
	out io.Writer
}

func NewColoredConsole(out io.Writer) coloredConsole {
	return coloredConsole{
		out: out,
	}
}

func (c coloredConsole) Print(color color.Color, args ...interface{}) error {
	_, err := color.Fprint(c.out, args...)

	if err != nil {
		return err
	}

	return nil
}

func (c coloredConsole) Println(color color.Color, args ...interface{}) error {
	_, err := color.Fprintln(c.out, args...)

	if err != nil {
		return err
	}

	return nil
}

func (c coloredConsole) Sprint(color color.Color, args ...interface{}) string {
	res := color.Sprint(args...)

	return res
}

func (c coloredConsole) Sprintln(color color.Color, args ...interface{}) string {
	res := color.Sprintln(args...)

	return res
}
