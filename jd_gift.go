package jd

type GiftService interface {
	// 礼金创建
	//    文档: https://union.jd.com/openplatform/api/12246
	CouponGiftGet(*CouponGiftGetRequest) (*CouponGiftGetResult, error)
	// 礼金停止
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftStop(*CouponGiftStopRequest) (*CouponGiftStopResult, error)
	// 礼金停止
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftStopBy(giftCouponKey string) (*CouponGiftStopResult, error)
	// 礼金效果数据
	//    文档: https://union.jd.com/openplatform/api/12248
	StatisticGiftCouponQuery(*StatisticGiftCouponQueryRequest) (*StatisticGiftCouponQueryResult, error)
}

type GiftServiceImpl struct {
	service Service
}

func newGiftService(service Service) GiftService {
	return &GiftServiceImpl{
		service: service,
	}
}

// 礼金创建
//    文档: https://union.jd.com/openplatform/api/12246
func (p *GiftServiceImpl) CouponGiftGet(request *CouponGiftGetRequest) (*CouponGiftGetResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftGetResult
	err := p.service.Do(&res, GiftGet, param)
	return &res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStop(request *CouponGiftStopRequest) (*CouponGiftStopResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftStopResult
	err := p.service.Do(&res, GiftStop, param)
	return &res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopBy(giftCouponKey string) (*CouponGiftStopResult, error) {
	return p.CouponGiftStop(&CouponGiftStopRequest{
		GiftCouponKey: giftCouponKey,
	})
}

// 礼金效果数据
//    文档: https://union.jd.com/openplatform/api/12248
func (p *GiftServiceImpl) StatisticGiftCouponQuery(request *StatisticGiftCouponQueryRequest) (*StatisticGiftCouponQueryResult, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res StatisticGiftCouponQueryResult
	err := p.service.Do(&res, GiftStatistic, param)
	return &res, err
}
