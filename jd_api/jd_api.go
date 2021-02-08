// Package jd_api provides JD_API Collection.
package jd_api

import (
	"github.com/cliod/jd-go"
	"github.com/cliod/jd-go/common"
)

// 京东联盟API.
//  文档: https://union.jd.com/openplatform/api.
type Service interface {
	// 仅用于继承HTTP API，实现通过组合方式
	common.Service

	// 获取配置
	GetConfig() *jd.Config
	// 设置配置
	SetConfig(config *jd.Config)
	// 设置http请求
	SetHttpService(common.Service)

	// 获取路由api
	GetRouteApi() string
	// 设置路由api
	SetRouteApi(string)

	/* 执行http请求并解析结果.
	   可通过 Execute 方法执行其他未封装的京东联盟API. */
	Execute(param jd.IParam) ([]byte, error)

	/* 执行http请求并解析结果.
	   参数:  v:     数据解析对象(指针).
	         param: 业务参数.
	   可通过 Request 方法执行其他未封装的京东联盟API.
	*/
	Request(v interface{}, param jd.IParam) error
}

type JdService struct {
	// 继承顶级接口
	Service
}
