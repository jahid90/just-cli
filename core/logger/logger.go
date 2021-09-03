package logger

import (
	"os"

	"github.com/fatih/color"
	"github.com/jahid90/just/config"
	"github.com/jahid90/just/output/console/colorize"
	"github.com/jahid90/just/output/console/plain"
)

func Debug(args ...interface{}) error {
	if config.LogLevel <= config.DEBUG {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, colorize.Sprint(*color.New(color.FgBlue), "[debug]"))
		printArgs = append(printArgs, args...)

		err := plain.Println(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil
}

func Info(args ...interface{}) error {
	if config.LogLevel <= config.INFO {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, colorize.Sprint(*color.New(color.FgGreen), "[info]"))
		printArgs = append(printArgs, args...)

		err := plain.Println(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil

}

func Warn(args ...interface{}) error {
	if config.LogLevel <= config.WARN {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, colorize.Sprint(*color.New(color.FgYellow), "[warn]"))
		printArgs = append(printArgs, args...)

		err := plain.Println(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil

}

func Error(args ...interface{}) error {
	if config.LogLevel <= config.ERROR {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, colorize.Sprint(*color.New(color.FgRed), "[error]"))
		printArgs = append(printArgs, args...)

		err := plain.Println(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil
}

func Fatal(args ...interface{}) {
	if config.LogLevel <= config.FATAL {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, colorize.Sprint(*color.New(color.FgHiRed), "[fatal]"))
		printArgs = append(printArgs, args...)

		plain.Println(printArgs...)
		os.Exit(1)
	}
}
