package logx

import (
	"fmt"
)

// Foreground text colors
const (
	fgBlack = uint8(iota + 30)
	fgRed
	fgGreen
	fgYellow
	fgBlue
	fgMagenta
	fgCyan
	fgWhite
)
const (
	color_red = uint8(iota + 91)
	color_green
	color_yellow
	color_blue
	color_magenta //洋红色
	info          = "[info]"
	trac          = "[trac]"
	erro          = "[error]"
	warn          = "[warn]"
	succ          = "[succ]"
)

func GreenPrintln(format string, a ...interface{}) {
	fmt.Println(green(fmt.Sprintf(format, a...)))
}
func RedPrintln(format string, a ...interface{}) {
	fmt.Println(red(fmt.Sprintf(format, a...)))
}
func BluePrintln(format string, a ...interface{}) {
	fmt.Println(blue(fmt.Sprintf(format, a...)))
}

func InfoWithGreen(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(prefix, green(fmt.Sprintf(format, a...)))
}
func InfoWithMagenta(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(prefix, magenta(fmt.Sprintf(format, a...)))
}
func InfoWithRed(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(prefix, red(fmt.Sprintf(format, a...)))
}
func InfoWithCyan(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(prefix, cyan(fmt.Sprintf(format, a...)))
}
func SuccessWithGreen(format string, a ...interface{}) {
	prefix := green(succ)
	fmt.Println(prefix, green(fmt.Sprintf(format, a...)))
}
func ErrorWithRed(format string, a ...interface{}) {
	prefix := red(erro)
	fmt.Println(prefix, red(fmt.Sprintf(format, a...)))
}
func WarningWithMagenta(format string, a ...interface{}) {
	prefix := magenta(succ)
	fmt.Println(prefix, magenta(fmt.Sprintf(format, a...)))
}
func TraceWithYellow(format string, a ...interface{}) {
	prefix := yellow(succ)
	fmt.Println(prefix, yellow(fmt.Sprintf(format, a...)))
}

func Trace(format string, a ...interface{}) {
	prefix := yellow(trac)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}
func Info(format string, a ...interface{}) {
	prefix := blue(info)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}

func Success(format string, a ...interface{}) {
	prefix := green(succ)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}
func Warning(format string, a ...interface{}) {
	prefix := magenta(warn)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}
func Error(format string, a ...interface{}) {
	prefix := red(erro)
	fmt.Println(prefix, fmt.Sprintf(format, a...))
}
func cyan(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", fgCyan, s)
}
func red(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_red, s)
}
func green(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_green, s)
}
func yellow(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_yellow, s)
}
func blue(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_blue, s)
}
func magenta(s string) string {
	return fmt.Sprintf("\x1b[%dm%s\x1b[0m", color_magenta, s)
}

//func formatLog(prefix string) string {
//	return time.Now().Format("2006/01/02 15:04:05") + " " + prefix + " "
//}
