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
)

const (
	Prefix = "[京东联盟]"
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

	trace *log.Logger
	debug *log.Logger
	error *log.Logger
	warn  *log.Logger
	info  *log.Logger
}

func NewLogger(level Level) Logger {
	return &LoggerImpl{
		level: level,
		trace: log.New(os.Stdout, "[TRACE] ", log.Ldate|log.Ltime|log.Lshortfile),
		debug: log.New(os.Stdout, "[DEBUG] ", log.Ldate|log.Ltime|log.Lshortfile),
		error: log.New(os.Stdout, "[ERROR] ", log.Ldate|log.Ltime|log.Lshortfile),
		warn:  log.New(os.Stdout, "[WARN] ", log.Ldate|log.Ltime|log.Lshortfile),
		info:  log.New(os.Stdout, "[INFO] ", log.Ldate|log.Ltime|log.Lshortfile),
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
		l.getLog(level).Println(colors[5](Prefix), colors[level](fmt.Sprintf(format, args...)))
	}
}

func (l *LoggerImpl) SetLevel(level Level) {
	l.level = level
}

func (l *LoggerImpl) getLog(level Level) *log.Logger {
	var lg *log.Logger
	switch level {
	case LevelDebug:
		lg = l.debug
	case LevelInfo:
		lg = l.info
	case LevelWarn:
		lg = l.warn
	case LevelError:
		lg = l.error
	case LevelTrace:
		lg = l.trace
	default:
		lg = l.warn
	}
	return lg
}

func (l *LoggerImpl) checkLevel(level Level) bool {
	return l.level > level
}
