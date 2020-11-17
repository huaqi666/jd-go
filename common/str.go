package common

import (
	"crypto/md5"
	"encoding/hex"
	"strings"
)

// Md5 签名
func Md5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

func Md5ToUpper(str string) string {
	return strings.ToUpper(Md5(str))
}
