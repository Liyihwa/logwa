package logwa

type MultipleLogger struct {
	logges []Logger
}

func NewMultipleLogger() Logger {
	return &MultipleLogger{}
}

func (ml *MultipleLogger) AddLogger(logger Logger) {
	ml.logges = append(ml.logges, logger)
}

func (ml *MultipleLogger) Debug(fmtString string, args ...any) {
	for _, l := range ml.logges {
		l.Debug(fmtString, args...)
	}
}

func (ml *MultipleLogger) Info(fmtString string, args ...any) {
	for _, l := range ml.logges {
		l.Info(fmtString, args...)
	}
}

func (ml *MultipleLogger) Warn(fmtString string, args ...any) {
	for _, l := range ml.logges {
		l.Warn(fmtString, args...)
	}
}

func (ml *MultipleLogger) Erro(fmtString string, args ...any) {
	for _, l := range ml.logges {
		l.Erro(fmtString, args...)
	}
}
