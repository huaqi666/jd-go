package log

import (
	"testing"
	"time"
)

func TestLoggerImpl_FileLog(t *testing.T) {
	// 将日志存入文件
	// 最后两个参数不能同时使用
	FileLog("logs/", "jd-go.log", 7, 1*time.Hour, 24*time.Minute)
}
