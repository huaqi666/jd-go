package jd

type OtherService interface {
	// 查询推广位【申请】
	PositionQuery(*PositionQueryRequest) (*PositionQueryResult, error)
	// 创建推广位【申请】
	PositionCreate(*PositionCreateRequest) (*PositionCreateResult, error)
	// 获取PID【申请】
	UserPidGet(*UserPidGetRequest) (*UserPidGetResult, error)
	// 活动查询接口
	ActivityQuery(*ActivityQueryRequest) (*ActivityQueryResult, error)
	// 京享红包效果数据
	StatisticsRedpacketQuery(*StatisticsRedpacketRequest) (*StatisticsRedpacketResult, error)
}

type OtherServiceImpl struct {
	service Service
}

func newOtherService(service Service) OtherService {
	return &OtherServiceImpl{
		service: service,
	}
}

func (o *OtherServiceImpl) PositionQuery(request *PositionQueryRequest) (*PositionQueryResult, error) {
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res PositionQueryResult
	err := o.service.Do(&res, PositionQuery, param)
	return &res, err
}

func (o *OtherServiceImpl) PositionCreate(request *PositionCreateRequest) (*PositionCreateResult, error) {
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res PositionCreateResult
	err := o.service.Do(&res, PositionCreate, param)
	return &res, err
}

func (o *OtherServiceImpl) UserPidGet(request *UserPidGetRequest) (*UserPidGetResult, error) {
	param := map[string]interface{}{}
	param["pidReq"] = request
	var res UserPidGetResult
	err := o.service.Do(&res, UserPidGet, param)
	return &res, err
}

func (o *OtherServiceImpl) ActivityQuery(request *ActivityQueryRequest) (*ActivityQueryResult, error) {
	param := map[string]interface{}{}
	param["activityReq"] = request
	var res ActivityQueryResult
	err := o.service.Do(&res, ActivityQuery, param)
	return &res, err
}

func (o *OtherServiceImpl) StatisticsRedpacketQuery(request *StatisticsRedpacketRequest) (*StatisticsRedpacketResult, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res StatisticsRedpacketResult
	err := o.service.Do(&res, StatisticsRedpacketQuery, param)
	return &res, err
}
