package logwa

import (
	"os"
)

/*
logwa是一个日志记录器,包含的功能如下:
1.	分级别的日志输出
2.	自定义Format格式
3.	自定义输出target,文件或者标准输出

定义日志格式时,需要指定Config结构体,Config字段如下:
UseColor 		bool: 			是否需要显示颜色
Target			io.Write: 		写入的位置
Level 			uint32:			日志等级,有DEBUG,INFO,WARN,ERRO四种
DataTimeFormat	string:			日志时间格式
LogMethods		[4]LogMethod:	分别对应四个等级日志记录时触发的函数

LogMethod定义如下:
type LogMethod func(datetime string, level string, message string) string
其中,datetime,level,message都会有框架提供,用户需要自己返回字符串,作为日志格式,例如:

	func defaultWarnMethod(datetime string, level string, message string) string {
		return fmt.Sprintf("%s {yx}%5s{;} : %s", datetime, level, message)
	}

其中的 {yx} {;}是cfmt中用到的转义字符
*/

// -----

var Std *Logger

func DefaultConfig() Config {
	return Config{
		UseColor:       true,
		Level:          INFO,
		Target:         os.Stdout,
		DataTimeFormat: DefaultDateTimeFormat,
		LogMethods:     [4]LogMethod{DefaultDebugMethod, DefaultInfoMethod, DefaultWarnMethod, DefaultErrorMethod},
	}
}

func init() {
	Std = NewLogger(DefaultConfig())
}

func Debug(fmt string, args ...any) {
	Std.Debug(fmt, args...)
}

func Info(fmt string, args ...any) {
	Std.Info(fmt, args...)
}
func Warn(fmt string, args ...any) {
	Std.Warning(fmt, args...)
}
func Erro(fmt string, args ...any) {
	Std.Error(fmt, args...)
}
