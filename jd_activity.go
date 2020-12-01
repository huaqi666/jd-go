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
func (o *ActivityServiceImpl) ActivityQuery(request *ActivityQueryRequest) (*ActivityQueryResult, error) {
	param := map[string]interface{}{}
	param["activityReq"] = request
	var res ActivityQueryResult
	err := o.service.Do(&res, ActivityQuery, param)
	return &res, err
}

// 京享红包效果数据
//    文档: https://union.jd.com/openplatform/api/14416
func (o *ActivityServiceImpl) StatisticsRedpacketQuery(request *StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res StatisticsRedpacketQueryResult
	err := o.service.Do(&res, StatisticsRedpacketQuery, param)
	return &res, err
}

// 活动查询接口
//    文档: https://union.jd.com/openplatform/api/12667
func (o *ActivityServiceImpl) ActivityQueryResult(request *ActivityQueryRequest) ([]byte, error) {
	res, err := o.ActivityQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 京享红包效果数据
//    文档: https://union.jd.com/openplatform/api/14416
func (o *ActivityServiceImpl) StatisticsRedpacketQueryResult(request *StatisticsRedpacketQueryRequest) ([]byte, error) {
	res, err := o.StatisticsRedpacketQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 活动查询接口
//    文档: https://union.jd.com/openplatform/api/12667
func (o *ActivityServiceImpl) ActivityQueryMap(request *ActivityQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["activityReq"] = request
	var res map[string]interface{}
	err := o.service.Do(&res, ActivityQuery, param)
	return res, err
}

// 京享红包效果数据
//    文档: https://union.jd.com/openplatform/api/14416
func (o *ActivityServiceImpl) StatisticsRedpacketQueryMap(request *StatisticsRedpacketQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res map[string]interface{}
	err := o.service.Do(&res, StatisticsRedpacketQuery, param)
	return res, err
}
