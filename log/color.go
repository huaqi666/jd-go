package log

/* 颜色
格式：\033[显示方式;前景色;背景色m

说明：
前景色            背景色           颜色
---------------------------------------
30                40              黑色
31                41              红色
32                42              绿色
33                43              黃色
34                44              蓝色
35                45              紫红色
36                46              青蓝色
37                47              白色
显示方式           意义
-------------------------
0                终端默认设置
1                高亮显示
4                使用下划线
5                闪烁
7                反白显示
8                不可见

例子：
\033[1;31;40m    <!--1-高亮显示 31-前景色红色  40-背景色黑色-->
\033[0m          <!--采用终端默认设置，即取消颜色设置-->
*/
type brush func(string) string

func newBrush(color string) brush {
	pre := "\033["
	reset := "\033[0m"
	return func(text string) string {
		return pre + color + "m" + text + reset
	}
}

//鉴于终端的通常使用习惯，一般白色和黑色字体是不可行的,所以30,37不可用，
var colors = []brush{
	//newBrush("1;41"), // Emergency          红色底
	//newBrush("1;35"), // Alert              紫色
	//newBrush("1;34"), // Critical           蓝色
	newBrush("1;31"), // Error              红色
	newBrush("1;33"), // Warn               黄色
	newBrush("1;36"), // Informational      天蓝色
	newBrush("1;32"), // Debug              绿色
	newBrush("1;37"), // Trace              白色
	newBrush("1;37"), // Empty              绿色
}
