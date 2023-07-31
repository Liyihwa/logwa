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

func (l *Logger) write(level string, method LogMethod, fmtString string, args ...any) {
	_, err := l.Target.Write([]byte(fmt.Sprintf(cfmt(l.UseColor, method(time.Now().Format(l.DataTimeFormat), level, fmtString)), args...)))
	if err != nil {
		panic(err)
	}
}

func (l *Logger) Debug(fmtString string, args ...any) {
	if l.Level <= DEBUG {
		l.write("DEBUG", l.LogMethods[DEBUG], fmtString, args...)
	}
}

func (l *Logger) Info(fmtString string, args ...any) {
	if l.Level <= INFO {
		l.write("INFO", l.LogMethods[INFO], fmtString, args...)
	}
}

func (l *Logger) Warn(fmtString string, args ...any) {
	if l.Level <= WARNING {
		l.write("WARN", l.LogMethods[WARNING], fmtString, args...)
	}
}

func (l *Logger) Erro(fmtString string, args ...any) {
	l.write("ERRO", l.LogMethods[ERROR], fmtString, args...)
	noColor := fmt.Sprintf(cfmt(false, fmtString), args...)
	panic(noColor)
}
