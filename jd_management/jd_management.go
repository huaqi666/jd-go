// Package jd_management provides management tools api of JD_API.
package jd_management

import "github.com/cliod/jd-go"

// 京东管理工具API(兼容)
type JdManagementService interface {
	// 创建推广位【申请】
	//   文档: https://union.jd.com/openplatform/api/10429
	JdPositionCreate(*jd.PositionCreateRequest) ([]byte, error)
	// 查询推广位【申请】
	//   文档: https://union.jd.com/openplatform/api/10428
	JdPositionQuery(*jd.PositionQueryRequest) ([]byte, error)
	// 获取PID【申请】
	// 此接口属于版本: 1.1
	//   文档: https://union.jd.com/openplatform/api/10430
	JdUserPidGet(*jd.UserPidGetRequest) ([]byte, error)
}

// 京东管理工具API(兼容)
type JdManagementServiceCompatible interface {
	JdManagementService
	// 创建推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreate(*jd.PositionCreateRequest) (*jd.PositionCreateResult, error)
	// 查询推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQuery(*jd.PositionQueryRequest) (*jd.PositionQueryResult, error)
	// 获取PID【申请】
	// 此接口属于版本: 1.1
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGet(*jd.UserPidGetRequest) (*jd.UserPidGetResult, error)
}
