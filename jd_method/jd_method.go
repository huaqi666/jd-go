// 接口路由方法
package jd_method

import "strings"

// 请求的路由方法(接口)
type Method string

func (m Method) String() string {
	return string(m)
}

func MethodOf(key ResponseKey) Method {
	l := len(key.String())
	return Method(strings.ReplaceAll(key.String()[:l-len("_response")], "_", "."))
}

// 响应的键名
type ResponseKey string

func (m ResponseKey) String() string {
	return string(m)
}

func KeyOfResponse(method Method) ResponseKey {
	return ResponseKey(strings.ReplaceAll(method.String(), ".", "_") + "_response")
}

func KeyOfResponce(method Method) ResponseKey {
	return ResponseKey(strings.ReplaceAll(method.String(), ".", "_") + "_responce")
}
