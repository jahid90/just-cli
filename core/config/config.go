package config

type LOG_LEVEL int

const (
	DEBUG LOG_LEVEL = iota
	INFO
	WARN
	ERROR
	FATAL
)

// Set to info by default; can be overridden
var LogLevel LOG_LEVEL = INFO
