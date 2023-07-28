package logger

import (
	"fmt"
	Level2 "github.com/Liyihwa/logwa/level"
	"io"
)

type Config struct {
	UseColor       bool
	Target         io.Writer
	Level          Level2.LogLevel
	DataTimeFormat string
	LogMethods     [4]LogMethod
}

const DefaultDateTimeFormat = "06-01-02 15:04:05.000"

func configVerif(config *Config) {
	if config == nil {
		panic("Config should not be nil")
	}
	if config.Level < Level2.DEBUG || config.Level > Level2.ERROR {
		panic("Config's Level out of range")
	}
	var i Level2.LogLevel
	for i = config.Level; i <= Level2.ERROR; i++ {
		if config.LogMethods[i] == nil {
			panic(fmt.Sprintf("Config's %s Method shouldn't be nil", config.Level.String()))
		}
	}
}
