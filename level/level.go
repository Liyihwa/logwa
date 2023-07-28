package level

import "strings"

type LogLevel int

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "DEBUG"
	case INFO:
		return "INFO"
	case WARNING:
		return "WARN"
	case ERROR:
		return "ERRO"
	}
	return ""
}

func LevelFromString(level string) LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
	case "warning":
		return WARNING
	case "erro":
	case "error":
		return ERROR
	}
	return -1
}
