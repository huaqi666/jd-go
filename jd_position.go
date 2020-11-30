package jd

type PositionService interface {
	// 查询推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQuery(*PositionQueryRequest) (*PositionQueryResult, error)
	// 创建推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreate(*PositionCreateRequest) (*PositionCreateResult, error)
	// 获取PID【申请】
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGet(*UserPidGetRequest) (*UserPidGetResult, error)
}

type PositionServiceImpl struct {
	service Service
}

func newPositionService(service Service) PositionService {
	return &PositionServiceImpl{
		service: service,
	}
}

// 查询推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10428
func (o *PositionServiceImpl) PositionQuery(request *PositionQueryRequest) (*PositionQueryResult, error) {
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res PositionQueryResult
	err := o.service.Do(&res, PositionQuery, param)
	return &res, err
}

// 创建推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10429
func (o *PositionServiceImpl) PositionCreate(request *PositionCreateRequest) (*PositionCreateResult, error) {
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res PositionCreateResult
	err := o.service.Do(&res, PositionCreate, param)
	return &res, err
}

// 获取PID【申请】
//    文档: https://union.jd.com/openplatform/api/10430
func (o *PositionServiceImpl) UserPidGet(request *UserPidGetRequest) (*UserPidGetResult, error) {
	param := map[string]interface{}{}
	param["pidReq"] = request
	var res UserPidGetResult
	err := o.service.Do(&res, UserPidGet, param)
	return &res, err
}
