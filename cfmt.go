package logger

import (
	"fmt"
	"strings"
)

/*
	cfmt 可以将带有format格式的字符串转码为带颜色的字符串
	例如格式化字符串:	I am a {b}boy{;}
	中的 boy会变为蓝色,其中{;}表示reset,关闭颜色
	再例如:	I am a {r_p}girl{;}
	中的 girl 会变为紫底红字

	{}中的字母会被解析为颜色, _字母为背景色, 单字母为字体色
	各种颜色如下:



		b 	黑
		r 	红
		g 	绿
		y 	黄
		u 	蓝
		p 	紫
		c 	青
		a 	灰
		x	使用粗体

	useColor bool:是否开启颜色,false时输出原字符串
	fmtStr string: 带有颜色的fmt字符串,例如 {u}%s abcv{;}
	args ... any : 被解析的参数
*/

var fontColors = []string{
	"30;",
	"31;",
	"32;",
	"33;",
	"34;",
	"35;",
	"36;",
	"37;",
	"1;",
}

var backColors = []string{
	"40;",
	"41;",
	"42;",
	"43;",
	"44;",
	"45;",
	"46;",
	"47;",
}

const (
	_BLACK  = 0
	_RED    = 1
	_GREEN  = 2
	_YELLOW = 3
	_BLUE   = 4
	_PURPLE = 5
	_CYAN   = 6
	_GRAY   = 7
	_BOLD   = 8
)

var colorMap = map[byte]int{
	'b': _BLACK,
	'r': _RED,
	'g': _GREEN,
	'y': _YELLOW,
	'u': _BLUE,
	'p': _PURPLE,
	'c': _CYAN,
	'a': _GRAY,
	'x': _BOLD,
}

const reset = "0"

func getColor(format string) string {
	sb := strings.Builder{}
	n := len(format)
	for i := 0; i < n; i++ {
		switch format[i] {
		case '_':
			if i+1 < n {
				if color, ok := colorMap[format[i+1]]; ok {
					sb.WriteString(backColors[color])
				}
			}
			i++
		case ';':
			sb.WriteString(reset)
		default:
			if color, ok := colorMap[format[i]]; ok {
				sb.WriteString(fontColors[color])
			}
		}
	}
	res := sb.String()
	if len(res) == 0 {
		return ""
	}
	return "\033[" + res[0:len(res)-1] + "m"
}
func cfmt(useColor bool, fmtStr string, args ...any) string {
	var res []byte
	n := len(fmtStr)

	for i := 0; i < n; i++ {
		switch fmtStr[i] {
		case '\\':
			continue
		case '{':
			j := i + 1
			for ; j < n && fmtStr[j] != '}'; j++ {
			}
			if useColor {
				res = append(res, getColor(fmtStr[i+1:j])...)
			}
			i = j
		default:
			res = append(res, fmtStr[i])
		}
	}

	if len(args) > 0 {
		return fmt.Sprintf(string(res), args...)
	} else {
		return string(res)
	}
}
