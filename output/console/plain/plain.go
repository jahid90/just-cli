package plain

import "fmt"

func Print(args ...interface{}) error {
	_, err := fmt.Print(args...)

	if err != nil {
		return err
	}

	return nil
}

func Println(args ...interface{}) error {
	_, err := fmt.Println(args...)

	if err != nil {
		return err
	}

	return nil
}

func Printf(format string, args ...interface{}) error {
	_, err := fmt.Printf(format, args...)

	if err != nil {
		return err
	}

	return nil
}

func Sprint(args ...interface{}) string {
	out := fmt.Sprint(args...)

	return out
}

func Sprintln(args ...interface{}) string {
	out := fmt.Sprintln(args...)

	return out
}

func Sprintf(format string, args ...interface{}) string {
	out := fmt.Sprintf(format, args...)

	return out
}
