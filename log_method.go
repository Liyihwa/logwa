package logwa

import "fmt"

type LogMethod func(datetime string, level string, message string) string

func DefaultDebugMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {g}%5s{;} : %s\n", datetime, level, message)
}

func DefaultInfoMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {u}%5s{;} : %s\n", datetime, level, message)
}

func DefaultWarnMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {yx}%5s{;} : %s\n", datetime, level, message)
}

func DefaultErrorMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {_rx}%5s{;} : %s\n", datetime, level, message)
}
