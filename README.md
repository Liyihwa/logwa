# Logger

logger是一款简单的日志框架,它包含的功能如下:

1.   定义了颜色转码格式,允许用户在控制台输出具有颜色的字符串
2.   支持划分等级打印日志
3.   支持用户自定义日志格式,自定义时间格式,自定义不同等级方法。

以下是一个简短的使用demo:

```go
func main() {
	logger.Debug("debug")
	logger.Info("debug")
	logger.Warn("debug")
	logger.Erro("debug")
}
```

### Config

logger中最核心的部分是Config,具体来说,每个Config结构体内容如下:

```go
type Config struct {
	UseColor       bool				//是否需要显示颜色,如果target为文件时,建议关闭该选项
	Target         io.Writer		// 写入的位置
	Level          LogLevel			//日志等级,有DEBUG,INFO,WARN,ERRO四种
	DataTimeFormat string			//日志时间格式
	LogMethods     [4]LogMethod	//分别对应四个等级日志记录时触发的函数
}
type LogMethod func(datetime string, level string, message string) string
```

其中**,datetime 为格式化后的时间字符串,level 为日志等级,message 为日志内容**,这些都会由框架提供**,**用户需要自己返回字符串**,**作为日志格式**,**例如:

```go
func DefaultWarnMethod(datetime string, level string, message string) string {
		return fmt.Sprintf("%s {yx}%5s{;} : %s", datetime, level, message)
}
```

其中的`{yx}`和`{;}`都是用于解析颜色的,我们稍后提到。

此外,logger提供了`DefaultStdoutConfig()`,当然你也可以完全自定义Config:

```go
func DefaultStdoutConfig() *Config {
	return &Config{
		UseColor:       true,
		Level:          INFO,
		Target:         os.Stdout,
		DataTimeFormat: DefaultDateTimeFormat,
		LogMethods:     [4]LogMethod{DefaultDebugMethod, DefaultInfoMethod, DefaultWarnMethod, DefaultErrorMethod},
	}
}
```

四个Default...Method如下:

```go
func DefaultDebugMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s %5s : %s", datetime, level, message)
}

func DefaultInfoMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {u}%5s{;} : %s", datetime, level, message)
}

func DefaultWarnMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {yx}%5s{;} : %s", datetime, level, message)
}

func DefaultErrorMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s {rx}%5s{;} : %s", datetime, level, message)
}
```

### 自定义日期格式

日期格式的定义与go语言内置 `time.Format()`完全相同:

**demo**:

```go
func main() {
	c := logger.DefaultStdoutConfig()
	c.Level = logger.DEBUG
	c.DataTimeFormat = "06-01-02 15:04:05"
	logger.StdoutOn(c)
	logger.Debug("debug")
	logger.Info("debug")
	logger.Warn("debug")
	logger.Erro("debug")
}
```

### 文件写入:

logger在默认情况下,只有一个标准输出:

```go
func init() {
	loggers = make([]*Config, 1)
	StdoutOn(DefaultStdoutConfig())
}
func StdoutOn(config *Config) {
	configVerif(config)
	if config.Target != os.Stdout {
		panic("Config's Target should be Stdout")
	}
	loggers[0] = config
}
```

为了添加写入的文件,我们可以为其添加一个Config:

```go

func warnMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s [{yx}%-5s{;}] %s", datetime, level, message)
}

func errorMethod(datetime string, level string, message string) string {
	return fmt.Sprintf("%s [{rx}%-5s{;}] %s", datetime, level, message)
}
func main() {
	file, err := os.OpenFile("t.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		panic(err.Error())
	}
	defer file.Close()
	c := logger.Config{UseColor: false, Target: file, Level: logger.WARNING, DataTimeFormat: logger.DefaultDateTimeFormat, LogMethods: [4]logger.LogMethod{nil, nil, warnMethod, errorMethod}}
	logger.AddConfig(&c)
   logger.Debug("Error at line {r}%d{;}", 23)
	logger.Info("debug")
	logger.Warn("debug")
	logger.Erro("debug")
}
```

### 彩色输出:

控制台彩色输出的源代码在`cfmt.go`中,`cfmt`可以将带有特定格式的字符串转码为带颜色的字符串
例如格式化字符串:`I am a {b}boy{;}`中的boy会变为蓝色,该字符串中{;}表示reset即关闭颜色。

再例如:`I am a {r_p}girl{;}`中的 girl 会变为紫底红字。

在`cfmt`中,`{}`中的字母会被解析为颜色, `_字母`为背景色,单字母为字体色。各种颜色如下:

| 字母 | 颜色     |
| ---- | -------- |
| b    | 黑色     |
| r    | 红色     |
| g    | 绿色     |
| y    | 黄色     |
| u    | 蓝色     |
| p    | 紫色     |
| c    | 青色     |
| a    | 灰色     |
| x    | 使用粗体 |

注意,`_x`无任何效果。





# Todo

1.   没有考虑并发的情况。

