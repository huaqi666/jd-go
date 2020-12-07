package log

import "time"

var (
	std Logger
)

func init() {
	std = NewLogger(LevelDebug)
}

func Panic(args ...interface{}) {
	std.Panic(args...)
}

func Fatal(args ...interface{}) {
	std.Fatal(args...)
}

func Error(args ...interface{}) {
	std.Error(args...)
}

func Warn(args ...interface{}) {
	std.Warn(args...)
}

func Info(args ...interface{}) {
	std.Info(args...)
}

func Debug(args ...interface{}) {
	std.Debug(args...)
}

func Trace(args ...interface{}) {
	std.Trace(args...)
}

func Print(args ...interface{}) {
	std.Print(args...)
}

// 日志
func Log(level Level, args ...interface{}) {
	std.Log(level, args...)
}

// 设置是否打印
func SetLevel(level Level) {
	std.SetLevel(level)
}

func FileLog(logPath string, logFileName string, maxNum uint, maxAge, rotationTime time.Duration) {
	std.FileLog(logPath, logFileName, maxNum, maxAge, rotationTime)
}
