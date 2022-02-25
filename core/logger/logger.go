package logger

import (
	"os"

	"github.com/fatih/color"
	"github.com/jahid90/just/core/config"
)

type LoggerFunc func(args ...interface{}) error
type FormatterFunc func(format string, args ...interface{}) string
type ColorizerFunc func(color color.Color, args ...interface{}) string

// Injection points with default no-op implementations
var Logger LoggerFunc = func(args ...interface{}) error { return nil }
var Formatter FormatterFunc = func(format string, args ...interface{}) string { return "" }
var Colorizer ColorizerFunc = func(color color.Color, args ...interface{}) string { return "" }

func Debug(args ...interface{}) error {
	if config.AppLogLevel <= config.DEBUG {
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
	if config.AppLogLevel <= config.INFO {
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
	if config.AppLogLevel <= config.WARN {
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
	if config.AppLogLevel <= config.ERROR {
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
	if config.AppLogLevel <= config.FATAL {
		var printArgs = make([]interface{}, 0)
		printArgs = append(printArgs, Colorizer(*color.New(color.FgHiRed), "[fatal]"))
		printArgs = append(printArgs, args...)

		Logger(printArgs...)
		os.Exit(1)
	}
}

func Debugf(format string, args ...interface{}) {
	if config.AppLogLevel <= config.DEBUG {
		formatted := Formatter(format, args...)
		Debug(formatted)
	}
}

func Infof(format string, args ...interface{}) {
	if config.AppLogLevel <= config.INFO {
		formatted := Formatter(format, args...)
		Info(formatted)
	}
}

func Warnf(format string, args ...interface{}) {
	if config.AppLogLevel <= config.WARN {
		formatted := Formatter(format, args...)
		Warn(formatted)
	}
}

func Errorf(format string, args ...interface{}) {
	if config.AppLogLevel <= config.ERROR {
		formatted := Formatter(format, args...)
		Error(formatted)
	}
}

func Fatalf(format string, args ...interface{}) {
	if config.AppLogLevel <= config.FATAL {
		formatted := Formatter(format, args...)
		Fatal(formatted)
	}
}
