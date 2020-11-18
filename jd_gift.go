package jd

type GiftService interface {
	// 礼金创建
	GitGet(*GiftGetRequest) (*GiftGetResult, error)
	// 礼金停止
	GiftStop(*GiftStopRequest) (*GiftStopResult, error)
	// 礼金效果数据
	GiftStatistic(*GiftStatisticRequest) (*GiftStatisticResult, error)
}

type GiftServiceImpl struct {
	service Service
}

func newGiftService(service Service) GiftService {
	return &GiftServiceImpl{
		service: service,
	}
}

func (p *GiftServiceImpl) GitGet(request *GiftGetRequest) (*GiftGetResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res GiftGetResult
	err := p.service.Do(&res, GiftGet, param)
	return &res, err
}

func (p *GiftServiceImpl) GiftStop(request *GiftStopRequest) (*GiftStopResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res GiftStopResult
	err := p.service.Do(&res, GiftStop, param)
	return &res, err
}

func (p *GiftServiceImpl) GiftStatistic(request *GiftStatisticRequest) (*GiftStatisticResult, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res GiftStatisticResult
	err := p.service.Do(&res, GiftStatistic, param)
	return &res, err
}
