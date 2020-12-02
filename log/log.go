package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Level int

const (
	LevelError Level = iota
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelEmpty
)

const (
	Prefix = "[京东联盟] "
)

var (
	levelStr = []string{
		colors[LevelError]("[ERROR] "),
		colors[LevelWarn]("[WARN] "),
		colors[LevelInfo]("[INFO] "),
		colors[LevelDebug]("[DEBUG] "),
		colors[LevelTrace]("[TRACE] "),
	}
)

// debug日志
type Logger interface {
	Trace(msg string, args ...interface{})
	Debug(msg string, args ...interface{})
	Info(msg string, args ...interface{})
	Warn(msg string, args ...interface{})
	Error(msg string, args ...interface{})
	// 日志
	Log(level Level, format string, args ...interface{})
	// 设置是否打印
	SetLevel(level Level)
}

type LoggerImpl struct {
	level Level

	output *log.Logger
}

func NewLogger(level Level) Logger {
	return &LoggerImpl{
		level:  level,
		output: log.New(os.Stdout, colors[LevelEmpty](Prefix), log.Ldate|log.Ltime|log.Lshortfile),
	}
}

func (l *LoggerImpl) Trace(msg string, args ...interface{}) {
	l.Log(LevelTrace, msg, args...)
}

func (l *LoggerImpl) Debug(msg string, args ...interface{}) {
	l.Log(LevelDebug, msg, args...)
}

func (l *LoggerImpl) Info(msg string, args ...interface{}) {
	l.Log(LevelInfo, msg, args...)
}

func (l *LoggerImpl) Warn(msg string, args ...interface{}) {
	l.Log(LevelWarn, msg, args...)
}

func (l *LoggerImpl) Error(msg string, args ...interface{}) {
	l.Log(LevelError, msg, args...)
}

func (l *LoggerImpl) Log(level Level, format string, args ...interface{}) {
	if l.checkLevel(level) {
		n := strings.Count(format, "%s")
		le := len(args)
		if !strings.HasSuffix(format, ":") {
			format += ":"
		}
		for i := 0; i < (le - n); i++ {
			format += " %s"
		}
		l.output.Println(levelStr[level], colors[level](fmt.Sprintf(format, args...)))
	}
}

func (l *LoggerImpl) SetLevel(level Level) {
	l.level = level
}

func (l *LoggerImpl) checkLevel(level Level) bool {
	return l.level >= level
}
