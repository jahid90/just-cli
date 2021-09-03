package logger

import (
	"os"

	"github.com/fatih/color"
	"github.com/jahid90/just/config"
)

type LoggerFunc func(args ...interface{}) error
type ColorizerFunc func(color color.Color, args ...interface{}) string

// Injection points
var Logger LoggerFunc = func(args ...interface{}) error { return nil }
var Colorizer ColorizerFunc = func(color color.Color, args ...interface{}) string { return "" }

func Debug(args ...interface{}) error {
	if config.LogLevel <= config.DEBUG {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, Colorizer(*color.New(color.FgBlue), "[debug]"))
		printArgs = append(printArgs, args...)

		err := Logger(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil
}

func Info(args ...interface{}) error {
	if config.LogLevel <= config.INFO {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, Colorizer(*color.New(color.FgGreen), "[info]"))
		printArgs = append(printArgs, args...)

		err := Logger(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil

}

func Warn(args ...interface{}) error {
	if config.LogLevel <= config.WARN {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, Colorizer(*color.New(color.FgYellow), "[warn]"))
		printArgs = append(printArgs, args...)

		err := Logger(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil

}

func Error(args ...interface{}) error {
	if config.LogLevel <= config.ERROR {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, Colorizer(*color.New(color.FgRed), "[error]"))
		printArgs = append(printArgs, args...)

		err := Logger(printArgs...)
		if err != nil {
			return err
		}
	}

	return nil
}

func Fatal(args ...interface{}) {
	if config.LogLevel <= config.FATAL {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, Colorizer(*color.New(color.FgHiRed), "[fatal]"))
		printArgs = append(printArgs, args...)

		Logger(printArgs...)
		os.Exit(1)
	}
}
