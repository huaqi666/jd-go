package log

import (
	"fmt"
	"log"
	"os"
	"strings"
)

type Level int

const (
	LevelPanic Level = iota
	LevelFatal
	LevelError
	LevelWarn
	LevelInfo
	LevelDebug
	LevelTrace
	LevelEmpty
)

const (
	Prefix = "[京东联盟] "
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
	l.Log(LevelTrace, "[TRACE] "+msg, args...)
}

func (l *LoggerImpl) Debug(msg string, args ...interface{}) {
	l.Log(LevelDebug, "[DEBUG] "+msg, args...)
}

func (l *LoggerImpl) Info(msg string, args ...interface{}) {
	l.Log(LevelInfo, "[INFO] "+msg, args...)
}

func (l *LoggerImpl) Warn(msg string, args ...interface{}) {
	l.Log(LevelWarn, "[WARN] "+msg, args...)
}

func (l *LoggerImpl) Error(msg string, args ...interface{}) {
	l.Log(LevelError, "[ERROR] "+msg, args...)
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
		l.output.Println(colors[level](fmt.Sprintf(format, args...)))
	}
}

func (l *LoggerImpl) SetLevel(level Level) {
	if level >= LevelEmpty || level < LevelError {
		level = LevelWarn
	}
	l.level = level
}

func (l *LoggerImpl) checkLevel(level Level) bool {
	return l.level >= level
}
