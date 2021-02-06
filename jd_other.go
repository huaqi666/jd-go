package jd

// Deprecated: 使用新接口: PositionService 和 ActivityService
type OtherService interface {
	// 查询推广位【申请】
	// Deprecated: 使用新接口: PositionService.PositionQuery
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQuery(*PositionQueryRequest) (*PositionQueryResult, error)
	// 创建推广位【申请】
	// Deprecated: 使用新接口: PositionService.PositionCreate
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreate(*PositionCreateRequest) (*PositionCreateResult, error)
	// 获取PID【申请】
	// Deprecated: 使用新接口: PositionService.UserPidGet
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGet(*UserPidGetRequest) (*UserPidGetResult, error)
	// 活动查询接口
	// Deprecated: 使用新接口: ActivityService.ActivityQuery
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQuery(*ActivityQueryRequest) (*ActivityQueryResult, error)
	// 京享红包效果数据
	// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQuery
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQuery(*StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error)

	// 查询推广位【申请】
	// Deprecated: 使用新接口: PositionService.PositionQueryResult
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQueryResult(*PositionQueryRequest) ([]byte, error)
	// 创建推广位【申请】
	// Deprecated: 使用新接口: PositionService.PositionCreateResult
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreateResult(*PositionCreateRequest) ([]byte, error)
	// 获取PID【申请】
	// Deprecated: 使用新接口: PositionService.UserPidGetResult
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGetResult(*UserPidGetRequest) ([]byte, error)
	// 活动查询接口
	// Deprecated: 使用新接口: ActivityService.ActivityQueryResult
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQueryResult(request *ActivityQueryRequest) ([]byte, error)
	// 京享红包效果数据
	// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQueryResult
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQueryResult(request *StatisticsRedpacketQueryRequest) ([]byte, error)

	// 查询推广位【申请】
	// Deprecated: 使用新接口: PositionService.PositionQueryMap
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQueryMap(*PositionQueryRequest) (map[string]interface{}, error)
	// 创建推广位【申请】
	// Deprecated: 使用新接口: PositionService.PositionCreateMap
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreateMap(*PositionCreateRequest) (map[string]interface{}, error)
	// 获取PID【申请】
	// Deprecated: 使用新接口: PositionService.UserPidGetMap
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGetMap(*UserPidGetRequest) (map[string]interface{}, error)
	// 活动查询接口
	// Deprecated: 使用新接口: ActivityService.ActivityQueryMap
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQueryMap(request *ActivityQueryRequest) (map[string]interface{}, error)
	// 京享红包效果数据
	// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQueryMap
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQueryMap(request *StatisticsRedpacketQueryRequest) (map[string]interface{}, error)
}

// Deprecated: 使用新接口: OtherServiceImpl 和 ActivityServiceImpl
type OtherServiceImpl struct {
	p PositionService
	a ActivityService
}

func newOtherService(p PositionService, a ActivityService) OtherService {
	return &OtherServiceImpl{
		p: p,
		a: a,
	}
}

// 查询推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionQuery
//    文档: https://union.jd.com/openplatform/api/10428
func (oth *OtherServiceImpl) PositionQuery(request *PositionQueryRequest) (*PositionQueryResult, error) {
	return oth.p.PositionQuery(request)
}

// 创建推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionCreate
//    文档: https://union.jd.com/openplatform/api/10429
func (oth *OtherServiceImpl) PositionCreate(request *PositionCreateRequest) (*PositionCreateResult, error) {
	return oth.p.PositionCreate(request)
}

// 获取PID【申请】
// Deprecated: 使用新接口: PositionService.UserPidGet
//    文档: https://union.jd.com/openplatform/api/10430
func (oth *OtherServiceImpl) UserPidGet(request *UserPidGetRequest) (*UserPidGetResult, error) {
	return oth.p.UserPidGet(request)
}

// 活动查询接口
// Deprecated: 使用新接口: ActivityService.ActivityQuery
//    文档: https://union.jd.com/openplatform/api/12667
func (oth *OtherServiceImpl) ActivityQuery(request *ActivityQueryRequest) (*ActivityQueryResult, error) {
	return oth.a.ActivityQuery(request)
}

// 京享红包效果数据
// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQuery
//    文档: https://union.jd.com/openplatform/api/14416
func (oth *OtherServiceImpl) StatisticsRedpacketQuery(request *StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error) {
	return oth.a.StatisticsRedpacketQuery(request)
}

// 查询推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionQueryResult
//    文档: https://union.jd.com/openplatform/api/10428
func (oth *OtherServiceImpl) PositionQueryResult(request *PositionQueryRequest) ([]byte, error) {
	return oth.p.PositionQueryResult(request)
}

// 创建推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionCreateResult
//    文档: https://union.jd.com/openplatform/api/10429
func (oth *OtherServiceImpl) PositionCreateResult(request *PositionCreateRequest) ([]byte, error) {
	return oth.p.PositionCreateResult(request)
}

// 获取PID【申请】
// Deprecated: 使用新接口: PositionService.UserPidGetResult
//    文档: https://union.jd.com/openplatform/api/10430
func (oth *OtherServiceImpl) UserPidGetResult(request *UserPidGetRequest) ([]byte, error) {
	return oth.p.UserPidGetResult(request)
}

// 活动查询接口
// Deprecated: 使用新接口: ActivityService.ActivityQueryResult
//    文档: https://union.jd.com/openplatform/api/12667
func (oth *OtherServiceImpl) ActivityQueryResult(request *ActivityQueryRequest) ([]byte, error) {
	return oth.a.ActivityQueryResult(request)
}

// 京享红包效果数据
// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQueryResult
//    文档: https://union.jd.com/openplatform/api/14416
func (oth *OtherServiceImpl) StatisticsRedpacketQueryResult(request *StatisticsRedpacketQueryRequest) ([]byte, error) {
	return oth.a.StatisticsRedpacketQueryResult(request)
}

// 查询推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionQueryMap
//    文档: https://union.jd.com/openplatform/api/10428
func (oth *OtherServiceImpl) PositionQueryMap(request *PositionQueryRequest) (map[string]interface{}, error) {
	return oth.p.PositionQueryMap(request)
}

// 创建推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionCreateMap
//    文档: https://union.jd.com/openplatform/api/10429
func (oth *OtherServiceImpl) PositionCreateMap(request *PositionCreateRequest) (map[string]interface{}, error) {
	return oth.p.PositionCreateMap(request)
}

// 获取PID【申请】
// Deprecated: 使用新接口: PositionService.UserPidGetMap
//    文档: https://union.jd.com/openplatform/api/10430
func (oth *OtherServiceImpl) UserPidGetMap(request *UserPidGetRequest) (map[string]interface{}, error) {
	return oth.p.UserPidGetMap(request)
}

// 活动查询接口
//
// Deprecated: 使用新接口: ActivityService.ActivityQueryMap
//    文档: https://union.jd.com/openplatform/api/12667
func (oth *OtherServiceImpl) ActivityQueryMap(request *ActivityQueryRequest) (map[string]interface{}, error) {
	return oth.a.ActivityQueryMap(request)
}

// 京享红包效果数据
// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQueryMap
//    文档: https://union.jd.com/openplatform/api/14416
func (oth *OtherServiceImpl) StatisticsRedpacketQueryMap(request *StatisticsRedpacketQueryRequest) (map[string]interface{}, error) {
	return oth.a.StatisticsRedpacketQueryMap(request)
}
