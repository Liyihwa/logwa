package logwa

type Logger interface {
	Debug(fmtString string, args ...any)
	Info(fmtString string, args ...any)
	Warn(fmtString string, args ...any)
	Erro(fmtString string, args ...any)
}
