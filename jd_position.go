package jd

type PositionService interface {
	// 查询推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQuery(*PositionQueryRequest) (*PositionQueryResult, error)
	// 创建推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreate(*PositionCreateRequest) (*PositionCreateResult, error)
	// 获取PID【申请】
	// 此接口属于版本: 1.1
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

	// 查询推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10428
	PositionQueryMap(*PositionQueryRequest) (map[string]interface{}, error)
	// 创建推广位【申请】
	//    文档: https://union.jd.com/openplatform/api/10429
	PositionCreateMap(*PositionCreateRequest) (map[string]interface{}, error)
	// 获取PID【申请】
	//    文档: https://union.jd.com/openplatform/api/10430
	UserPidGetMap(*UserPidGetRequest) (map[string]interface{}, error)
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
func (pos *PositionServiceImpl) PositionQuery(request *PositionQueryRequest) (*PositionQueryResult, error) {
	err := pos.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res PositionQueryResult
	err = pos.service.Do(&res, PositionQuery, param)
	return &res, err
}

// 创建推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10429
func (pos *PositionServiceImpl) PositionCreate(request *PositionCreateRequest) (*PositionCreateResult, error) {
	err := pos.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res PositionCreateResult
	err = pos.service.Do(&res, PositionCreate, param)
	return &res, err
}

// 获取PID【申请】
// 此接口在版本1.1才有
//    文档: https://union.jd.com/openplatform/api/10430
func (pos *PositionServiceImpl) UserPidGet(request *UserPidGetRequest) (*UserPidGetResult, error) {
	err := pos.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["pidReq"] = request
	var res UserPidGetResult
	err = pos.service.Do(&res, UserPidGet, param)
	return &res, err
}

// 查询推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10428
func (pos *PositionServiceImpl) PositionQueryResult(request *PositionQueryRequest) ([]byte, error) {
	res, err := pos.PositionQuery(request)
	return pos.service.GetResult(res, err)
}

// 创建推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10429
func (pos *PositionServiceImpl) PositionCreateResult(request *PositionCreateRequest) ([]byte, error) {
	res, err := pos.PositionCreate(request)
	return pos.service.GetResult(res, err)
}

// 获取PID【申请】
//    文档: https://union.jd.com/openplatform/api/10430
func (pos *PositionServiceImpl) UserPidGetResult(request *UserPidGetRequest) ([]byte, error) {
	res, err := pos.UserPidGet(request)
	return pos.service.GetResult(res, err)
}

// 查询推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10428
func (pos *PositionServiceImpl) PositionQueryMap(request *PositionQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res map[string]interface{}
	err := pos.service.Do(&res, PositionQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if pos.service.GetRouteApi() == JosRootEndpoint {
		key = PositionQueryResponce
	} else {
		key = PositionQueryResponse
	}
	return pos.service.ParseResult(res, key)
}

// 创建推广位【申请】
//    文档: https://union.jd.com/openplatform/api/10429
func (pos *PositionServiceImpl) PositionCreateMap(request *PositionCreateRequest) (map[string]interface{}, error) {
	err := pos.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["positionReq"] = request
	var res map[string]interface{}
	err = pos.service.Do(&res, PositionCreate, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if pos.service.GetRouteApi() == JosRootEndpoint {
		key = PositionCreateResponce
	} else {
		key = PositionCreateResponse
	}
	return pos.service.ParseResult(res, key)
}

// 获取PID【申请】
//    文档: https://union.jd.com/openplatform/api/10430
func (pos *PositionServiceImpl) UserPidGetMap(request *UserPidGetRequest) (map[string]interface{}, error) {
	err := pos.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["pidReq"] = request
	var res map[string]interface{}
	err = pos.service.Do(&res, UserPidGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if pos.service.GetRouteApi() == JosRootEndpoint {
		key = UserPidGetResponce
	} else {
		key = UserPidGetResponse
	}
	return pos.service.ParseResult(res, key)
}
