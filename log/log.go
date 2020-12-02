package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
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
	Panic(args ...interface{})
	Fatal(args ...interface{})
	Error(args ...interface{})
	Warn(args ...interface{})
	Info(args ...interface{})
	Debug(args ...interface{})
	Trace(args ...interface{})
	Print(args ...interface{})

	// 日志
	Logf(level Level, format string, args ...interface{})
	// 日志
	Log(level Level, args ...interface{})
	// 设置是否打印
	SetLevel(level Level)
}

type LoggerImpl struct {
	*logrus.Logger
}

func NewLogger(level Level) Logger {
	lgr := logrus.New()
	lgr.SetFormatter(newFormatter())
	lgr.Level = logrus.Level(level)
	return &LoggerImpl{
		Logger: lgr,
	}
}

func (l *LoggerImpl) Panic(args ...interface{}) {
	l.Log(LevelPanic, args...)
}

func (l *LoggerImpl) Fatal(args ...interface{}) {
	l.Log(LevelFatal, args...)
}

func (l *LoggerImpl) Error(args ...interface{}) {
	l.Log(LevelError, args...)
}

func (l *LoggerImpl) Warn(args ...interface{}) {
	l.Log(LevelWarn, args...)
}

func (l *LoggerImpl) Info(args ...interface{}) {
	l.Log(LevelInfo, args...)
}

func (l *LoggerImpl) Debug(args ...interface{}) {
	l.Log(LevelDebug, args...)
}

func (l *LoggerImpl) Trace(args ...interface{}) {
	l.Log(LevelTrace, args...)
}

func (l *LoggerImpl) Logf(level Level, format string, args ...interface{}) {
	if l.checkLevel(level) {
		n := strings.Count(format, "%s")
		le := len(args)
		for i := 0; i < (le - n); i++ {
			format += "%s "
		}
		l.Logger.Log(logrus.Level(level), fmt.Sprintf(format, args...))
	}
}

func (l *LoggerImpl) Log(level Level, args ...interface{}) {
	l.Logf(level, Prefix, args...)
}

func (l *LoggerImpl) SetLevel(level Level) {
	if level >= LevelEmpty || level < LevelPanic {
		level = LevelWarn
	}
	l.Level = logrus.Level(level)
}

func (l *LoggerImpl) checkLevel(level Level) bool {
	return uint64(l.Level) >= uint64(level)
}
