package logwa

import "fmt"

type LogMethod func(datetime string, level string, message string) string

// -------------------

func DefaultMethods() [4]LogMethod {
	return [4]LogMethod{DefaultDebug, DefaultInfo, DefaultWarn, DefaultError}
}

func DefaultDebug(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {g}%5s{;} : %s\n", datetime, level, message)
}

func DefaultInfo(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {u}%5s{;} : %s\n", datetime, level, message)
}

func DefaultWarn(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {yx}%5s{;} : %s\n", datetime, level, message)
}

func DefaultError(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {_rx}%5s{;} : %s\n", datetime, level, message)
}

// -------------------

func LevelOnlyDebug(datetime string, level string, message string) string {
	return fmt.Sprintf("{g}%5s{;} : %s\n", level, message)
}

func LevelOnlyInfo(datetime string, level string, message string) string {
	return fmt.Sprintf("{u}%5s{;} : %s\n", level, message)
}

func LevelOnlyWarn(datetime string, level string, message string) string {
	return fmt.Sprintf("{yx}%5s{;} : %s\n", level, message)
}

func LevelOnlyErro(datetime string, level string, message string) string {
	return fmt.Sprintf("{_rx}%5s{;} : %s\n", level, message)
}

func LevelOnlyMethods() [4]LogMethod {
	return [4]LogMethod{LevelOnlyDebug, LevelOnlyInfo, LevelOnlyWarn, LevelOnlyErro}
}
