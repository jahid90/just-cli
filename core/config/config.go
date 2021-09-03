package config

type LOG_LEVEL int

const (
	DEBUG LOG_LEVEL = iota
	INFO
	WARN
	ERROR
	FATAL
)

var LogLevel LOG_LEVEL = DEBUG

func SetLogLevel(environment string) {
	if environment == "production" {
		LogLevel = WARN
	} else {
		LogLevel = INFO
	}
}
