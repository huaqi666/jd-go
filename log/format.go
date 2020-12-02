package log

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"strings"
	"time"
)

// gt log formatter
type formatter struct{}

func (s *formatter) Format(entry *logrus.Entry) ([]byte, error) {
	timestamp := time.Now().Local().Format("2006-01-02 15:04:05")
	msg := fmt.Sprintf("\u001B[33m[%s]\u001B[0m \u001B[36;1m[%s]\u001B[0m %s\n",
		timestamp, strings.ToUpper(entry.Level.String()), colors[entry.Level](entry.Message))
	return []byte(msg), nil
}

func newFormatter() *formatter {
	return &formatter{}
}
