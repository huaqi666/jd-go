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
}

// Deprecated: 使用新接口: PositionServiceImpl 和 ActivityServiceImpl
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
func (o *OtherServiceImpl) PositionQuery(request *PositionQueryRequest) (*PositionQueryResult, error) {
	return o.p.PositionQuery(request)
}

// 创建推广位【申请】
// Deprecated: 使用新接口: PositionService.PositionCreate
//    文档: https://union.jd.com/openplatform/api/10429
func (o *OtherServiceImpl) PositionCreate(request *PositionCreateRequest) (*PositionCreateResult, error) {
	return o.p.PositionCreate(request)
}

// 获取PID【申请】
// Deprecated: 使用新接口: PositionService.UserPidGet
//    文档: https://union.jd.com/openplatform/api/10430
func (o *OtherServiceImpl) UserPidGet(request *UserPidGetRequest) (*UserPidGetResult, error) {
	return o.p.UserPidGet(request)
}

// 活动查询接口
// Deprecated: 使用新接口: ActivityService.ActivityQuery
//    文档: https://union.jd.com/openplatform/api/12667
func (o *OtherServiceImpl) ActivityQuery(request *ActivityQueryRequest) (*ActivityQueryResult, error) {
	return o.a.ActivityQuery(request)
}

// 京享红包效果数据
// Deprecated: 使用新接口: ActivityService.StatisticsRedpacketQuery
//    文档: https://union.jd.com/openplatform/api/14416
func (o *OtherServiceImpl) StatisticsRedpacketQuery(request *StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error) {
	return o.a.StatisticsRedpacketQuery(request)
}
