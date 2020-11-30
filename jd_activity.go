package jd

type ActivityService interface {
	// 活动查询接口
	//    文档: https://union.jd.com/openplatform/api/12667
	ActivityQuery(*ActivityQueryRequest) (*ActivityQueryResult, error)
	// 京享红包效果数据
	//    文档: https://union.jd.com/openplatform/api/14416
	StatisticsRedpacketQuery(*StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error)
}

type ActivityServiceImpl struct {
	service Service
}

func newActivityService(service Service) ActivityService {
	return &ActivityServiceImpl{
		service: service,
	}
}

func (o *ActivityServiceImpl) ActivityQuery(request *ActivityQueryRequest) (*ActivityQueryResult, error) {
	param := map[string]interface{}{}
	param["activityReq"] = request
	var res ActivityQueryResult
	err := o.service.Do(&res, ActivityQuery, param)
	return &res, err
}

func (o *ActivityServiceImpl) StatisticsRedpacketQuery(request *StatisticsRedpacketQueryRequest) (*StatisticsRedpacketQueryResult, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res StatisticsRedpacketQueryResult
	err := o.service.Do(&res, StatisticsRedpacketQuery, param)
	return &res, err
}
