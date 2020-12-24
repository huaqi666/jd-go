package jd

import "github.com/cliod/jd-go/common/cmap"

type IParam interface {
	// 请求的方法
	Method() string
	// 请求的参数
	Params() cmap.CMap
}

// 用户兼容
type TransferParam struct {
	method Method
	param  cmap.CMap
}

func NewTParam(method Method, v interface{}) IParam {
	return &TransferParam{
		method: method,
		param:  cmap.Struct(v),
	}
}

func (m TransferParam) Method() string {
	return m.method.String()
}

func (m TransferParam) Params() cmap.CMap {
	return m.param
}
