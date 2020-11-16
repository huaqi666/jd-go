package common

import (
	"crypto/md5"
	"encoding/hex"
	"sort"
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

func BuildSignStr(request interface{}, sk string, ignoreParams ...string) string {

	var params map[string]interface{}
	params, ok := request.(map[string]interface{})
	if !ok {
		params = ToMap(request)
	}
	if params == nil {
		return ""
	}

	var arr []string
	for k := range params {
		arr = append(arr, k)
	}
	sort.Strings(arr)
	ign := strings.Join(ignoreParams, ",")

	var sign string

	for _, key := range arr {
		shouldSign := false
		value := params[key]
		var vv = value.(string)
		if key != "" && vv != "" && !strings.Contains(ign, key) {
			shouldSign = true
		}
		if shouldSign {
			sign += key + vv
		}
	}
	return sk + sign + sk
}

func Sign(request interface{}, sk string, ignoreParams ...string) string {
	signStr := BuildSignStr(request, sk, ignoreParams...)
	return Md5ToUpper(signStr)
}
