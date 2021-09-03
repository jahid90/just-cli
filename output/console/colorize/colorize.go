package colorize

import "github.com/fatih/color"

func Print(color color.Color, args ...interface{}) error {
	_, err := color.Print(args...)

	if err != nil {
		return err
	}

	return nil
}

func Println(color color.Color, args ...interface{}) error {
	_, err := color.Println(args...)

	if err != nil {
		return err
	}

	return nil
}

func Sprint(color color.Color, args ...interface{}) string {
	out := color.Sprint(args...)

	return out
}

func Sprintln(color color.Color, args ...interface{}) string {
	out := color.Sprintln(args...)

	return out
}
