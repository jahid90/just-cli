package config

type LogLevel int

const (
	DEBUG LogLevel = iota
	INFO
	WARN
	ERROR
	FATAL
)

var AppLogLevel LogLevel = WARN

func SetLogLevel(logLevel string) {

	switch logLevel {
	case "DEBUG":
		AppLogLevel = DEBUG
	case "INFO":
		AppLogLevel = INFO
	case "WARN":
		AppLogLevel = WARN
	case "ERROR":
		AppLogLevel = ERROR
	case "FATAL":
		AppLogLevel = FATAL
	default:
		AppLogLevel = WARN
	}

}
