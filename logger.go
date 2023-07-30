package logwa

import (
	"fmt"
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
	if l.Level <= DEBUG {
		l.write("DEBUG", l.LogMethods[DEBUG], fmt.Sprintf(fmtString, args...))
	}
}

func (l *Logger) Info(fmtString string, args ...any) {
	if l.Level <= INFO {
		l.write("INFO", l.LogMethods[INFO], fmt.Sprintf(fmtString, args...))
	}
}

func (l *Logger) Warn(fmtString string, args ...any) {
	if l.Level <= WARNING {
		l.write("WARN", l.LogMethods[WARNING], fmt.Sprintf(fmtString, args...))
	}
}

func (l *Logger) Erro(fmtString string, args ...any) {
	fmtStr := fmt.Sprintf(fmtString, args...)
	l.write("ERRO", l.LogMethods[ERROR], fmtStr)
	panic(fmtStr)
}
