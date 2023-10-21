package logwa

/*
std log内定义了一个标准输出的Logger,开箱即用
*/
var Std *BaseLogger

func init() {
	Std = NewBaseLogger(DefaultConfig())
}
