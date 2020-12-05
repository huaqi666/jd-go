package jd

type ActivityService interface {
	// 活动查询接口
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQuery(*ActivityQueryRequest) (*ActivityQueryResult, error)
	// 京享红包效果数据
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQuery(*StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error)

	// 活动查询接口
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQueryResult(request *ActivityQueryRequest) ([]byte, error)
	// 京享红包效果数据
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQueryResult(request *StatisticsRedpacketQueryRequest) ([]byte, error)

	// 活动查询接口
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQueryMap(request *ActivityQueryRequest) (map[string]interface{}, error)
	// 京享红包效果数据
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQueryMap(request *StatisticsRedpacketQueryRequest) (map[string]interface{}, error)
}

type ActivityServiceImpl struct {
	service Service
}

func newActivityService(service Service) ActivityService {
	return &ActivityServiceImpl{
		service: service,
	}
}

// 活动查询接口
//    文档: https://union.jd.com/openplatform/api/12667
func (act *ActivityServiceImpl) ActivityQuery(request *ActivityQueryRequest) (*ActivityQueryResult, error) {
	err := act.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["activityReq"] = request
	var res ActivityQueryResult
	err = act.service.Do(&res, ActivityQuery, param)
	return &res, err
}

// 京享红包效果数据
//    文档: https://union.jd.com/openplatform/api/14416
func (act *ActivityServiceImpl) StatisticsRedpacketQuery(request *StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error) {
	err := act.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res StatisticsRedpacketQueryResult
	err = act.service.Do(&res, StatisticsRedpacketQuery, param)
	return &res, err
}

// 活动查询接口
//    文档: https://union.jd.com/openplatform/api/12667
func (act *ActivityServiceImpl) ActivityQueryResult(request *ActivityQueryRequest) ([]byte, error) {
	return act.service.GetResult(act.ActivityQuery(request))
}

// 京享红包效果数据
//    文档: https://union.jd.com/openplatform/api/14416
func (act *ActivityServiceImpl) StatisticsRedpacketQueryResult(request *StatisticsRedpacketQueryRequest) ([]byte, error) {
	return act.service.GetResult(act.StatisticsRedpacketQuery(request))
}

// 活动查询接口
//    文档: https://union.jd.com/openplatform/api/12667
func (act *ActivityServiceImpl) ActivityQueryMap(request *ActivityQueryRequest) (map[string]interface{}, error) {
	err := act.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["activityReq"] = request
	var res map[string]interface{}
	err = act.service.Do(&res, ActivityQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if act.service.GetRouteApi() == JosRootEndpoint {
		key = ActivityQueryResponce
	} else {
		key = ActivityQueryResponse
	}
	return act.service.ParseResult(res, key)
}

// 京享红包效果数据
//    文档: https://union.jd.com/openplatform/api/14416
func (act *ActivityServiceImpl) StatisticsRedpacketQueryMap(request *StatisticsRedpacketQueryRequest) (map[string]interface{}, error) {
	err := act.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res map[string]interface{}
	err = act.service.Do(&res, StatisticsRedpacketQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if act.service.GetRouteApi() == JosRootEndpoint {
		key = StatisticsRedpacketQueryResponce
	} else {
		key = StatisticsRedpacketQueryResponse
	}
	return act.service.ParseResult(res, key)
}
