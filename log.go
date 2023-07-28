package logger

import (
	"fmt"
	"io"
	"os"
	"strings"
	"time"
)

/*
	logger是一个日志记录器,包含的功能如下:
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

type Config struct {
	UseColor       bool
	Target         io.Writer
	Level          LogLevel
	DataTimeFormat string
	LogMethods     [4]LogMethod
}

func (c *Config) write(level string, method LogMethod, format string, args ...any) {
	_, err := c.Target.Write([]byte(cfmt(c.UseColor, method(time.Now().Format(c.DataTimeFormat), level, fmt.Sprintf(format, args...)))))
	if err != nil {
		panic(err.Error())
	}
}

type LogLevel int

func LevelFromString(level string) LogLevel {
	switch strings.ToLower(level) {
	case "debug":
		return DEBUG
	case "info":
		return INFO
	case "warn":
	case "warning":
		return WARNING
	case "erro":
	case "error":
		return ERROR
	}
	return -1
}

const (
	DEBUG = iota
	INFO
	WARNING
	ERROR
)

func (l LogLevel) String() string {
	switch l {
	case DEBUG:
		return "debug"
	case INFO:
		return "info"
	case WARNING:
		return "warn"
	case ERROR:
		return "erro"
	}
	panic(fmt.Sprintf("Level should be in range[0,3], not %d ", l))
}

const DefaultDateTimeFormat = "06-01-02 15:04:05.000"

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

//-----

var loggers []*Config
var stdCofig *Config

func init() {
	loggers = make([]*Config, 0)
	StdoutOn(DefaultStdoutConfig())
}

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

// StdoutOn 会为Logger定义标准输出
func StdoutOn(config *Config) {
	configVerif(config)
	stdCofig = config
}

func StdoutOff() {
	stdCofig = nil
}

func AddConfig(config *Config) {
	configVerif(config)
	if config.Target == os.Stdout {
		panic("Config's Target shouldn't be Stdout")
	}
	if config.Target == nil {
		panic("Config's Target shouldn't be nil")
	}

	loggers = append(loggers, config)
}

func DefaultStdoutConfig() *Config {
	return &Config{
		UseColor:       true,
		Level:          INFO,
		Target:         os.Stdout,
		DataTimeFormat: DefaultDateTimeFormat,
		LogMethods:     [4]LogMethod{DefaultDebugMethod, DefaultInfoMethod, DefaultWarnMethod, DefaultErrorMethod},
	}
}

func Debug(fmt string, args ...any) {
	if stdCofig != nil {
		stdCofig.write("DEBUG", stdCofig.LogMethods[DEBUG], fmt, args...)
	}
	for _, c := range loggers {
		if c.Level <= DEBUG {
			c.write("DEBUG", c.LogMethods[DEBUG], fmt, args...)
		}
	}
}

func Info(fmt string, args ...any) {
	if stdCofig != nil {
		stdCofig.write("INFO", stdCofig.LogMethods[INFO], fmt, args...)
	}
	for _, c := range loggers {
		if c.Level <= INFO {
			c.write("INFO", c.LogMethods[INFO], fmt, args...)
		}
	}
}
func Warn(fmt string, args ...any) {
	if stdCofig != nil {
		stdCofig.write("WARN", stdCofig.LogMethods[WARNING], fmt, args...)
	}
	for _, c := range loggers {
		if c.Level <= WARNING {
			c.write("WARN", c.LogMethods[WARNING], fmt, args...)
		}
	}
}
func Erro(fmt string, args ...any) {
	if stdCofig != nil {
		stdCofig.write("ERRO", stdCofig.LogMethods[ERROR], fmt, args...)
	}
	for _, c := range loggers {
		if c.Level <= ERROR {
			c.write("ERRO", c.LogMethods[ERROR], fmt, args...)
		}
	}
}
