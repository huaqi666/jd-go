package jd

type GiftService interface {
	// 礼金创建
	CouponGiftGet(*CouponGiftGetRequest) (*CouponGiftGetResult, error)
	// 礼金停止
	CouponGiftStop(*CouponGiftStopRequest) (*CouponGiftStopResult, error)
	// 礼金效果数据
	GiftStatisticCouponQuery(*GiftStatisticCouponQueryRequest) (*GiftStatisticCouponQueryResult, error)
}

type GiftServiceImpl struct {
	service Service
}

func newGiftService(service Service) GiftService {
	return &GiftServiceImpl{
		service: service,
	}
}

func (p *GiftServiceImpl) CouponGiftGet(request *CouponGiftGetRequest) (*CouponGiftGetResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftGetResult
	err := p.service.Do(&res, GiftGet, param)
	return &res, err
}

func (p *GiftServiceImpl) CouponGiftStop(request *CouponGiftStopRequest) (*CouponGiftStopResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftStopResult
	err := p.service.Do(&res, GiftStop, param)
	return &res, err
}

func (p *GiftServiceImpl) GiftStatisticCouponQuery(request *GiftStatisticCouponQueryRequest) (*GiftStatisticCouponQueryResult, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res GiftStatisticCouponQueryResult
	err := p.service.Do(&res, GiftStatistic, param)
	return &res, err
}
