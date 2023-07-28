package logger

import (
	"fmt"
	"github.com/Liyihwa/logwa/level"
	"time"
)

type Logger struct {
	Config
}

func NewLogger(c Config) *Logger {
	return &Logger{c}
}

func (l *Logger) write(level string, method LogMethod, message string) {
	var err error

	if l.UseColor {
		_, err = l.Target.Write([]byte(cfmt(method(time.Now().Format(l.DataTimeFormat), level, message))))
	} else {
		_, err = l.Target.Write([]byte(fmt.Sprintf(method(time.Now().Format(l.DataTimeFormat), level, message))))
	}

	if err != nil {
		panic(err.Error())
	}
}

func (l *Logger) Debug(fmtString string, args ...any) {
	l.write("DEBUG", l.LogMethods[level.DEBUG], fmt.Sprintf(fmtString, args...))
}

func (l *Logger) Info(fmtString string, args ...any) {
	l.write("INFO", l.LogMethods[level.INFO], fmt.Sprintf(fmtString, args...))
}

func (l *Logger) Warning(fmtString string, args ...any) {
	l.write("WARN", l.LogMethods[level.WARNING], fmt.Sprintf(fmtString, args...))
}

func (l *Logger) Error(fmtString string, args ...any) {
	l.write("ERRO", l.LogMethods[level.ERROR], fmt.Sprintf(fmtString, args...))
}
