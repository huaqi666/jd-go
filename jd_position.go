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
	// 查询推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQueryResult(*PositionQueryRequest) ([]byte, error)
	// 创建推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreateResult(*PositionCreateRequest) ([]byte, error)
	// 获取PID【申请】
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGetResult(*UserPidGetRequest) ([]byte, error)
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

// 查询推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10428
func (o *PositionServiceImpl) PositionQueryResult(request *PositionQueryRequest) ([]byte, error) {
	res, err := o.PositionQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 创建推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10429
func (o *PositionServiceImpl) PositionCreateResult(request *PositionCreateRequest) ([]byte, error) {
	res, err := o.PositionCreate(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 获取PID【申请】
//    文档: https://union.jd.com/openplatform/api/10430
func (o *PositionServiceImpl) UserPidGetResult(request *UserPidGetRequest) ([]byte, error) {
	res, err := o.UserPidGet(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}
