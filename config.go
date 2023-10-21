package logwa

import (
	"fmt"
	"io"
	"os"
)

type Config struct {
	UseColor       bool
	Target         io.Writer
	Level          LogLevel
	DataTimeFormat string
	LogMethods     [4]LogMethod
}

const DefaultDateTimeFormat = "06-01-02 15:04:05.000"

func configVerif(config *Config) {
	if config == nil {
		panic("Config should not be nil")
	}
	if config.Level < DEBUG || config.Level > ERROR {
		panic("Config's Level out of range")
	}
	var i LogLevel
	for i = config.Level; i <= ERROR; i++ {
		if config.LogMethods[i] == nil {
			panic(fmt.Sprintf("Config's %s Method shouldn't be nil", config.Level.String()))
		}
	}
}

func DefaultConfig() Config {
	return Config{
		UseColor:       true,
		Level:          INFO,
		Target:         os.Stdout,
		DataTimeFormat: DefaultDateTimeFormat,
		LogMethods:     DefaultMethods()}
}
