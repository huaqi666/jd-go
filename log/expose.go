package log

var (
	std = NewLogger(LevelDebug)
)

func Debug(msg string, args ...interface{}) {
	std.Debug(msg, args...)
}

func Info(msg string, args ...interface{}) {
	std.Info(msg, args...)
}

func Warn(msg string, args ...interface{}) {
	std.Warn(msg, args...)
}

func Error(msg string, args ...interface{}) {
	std.Error(msg, args...)
}

// 日志
func Log(level Level, format string, args ...interface{}) {
	std.Log(level, format, args...)
}

// 设置是否打印
func SetLevel(level Level) {
	std.SetLevel(level)
}
