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
	CouponGiftStopByKey(giftCouponKey string) (*CouponGiftStopResult, error)
	// 礼金效果数据
	//    文档: https://union.jd.com/openplatform/api/12248
	StatisticGiftCouponQuery(*StatisticGiftCouponQueryRequest) (*StatisticGiftCouponQueryResult, error)
	// 礼金停止
	// Deprecated: 使用 CouponGiftStopByKey
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftStopBy(giftCouponKey string) (*CouponGiftStopResult, error)

	// 礼金创建
	//    文档: https://union.jd.com/openplatform/api/12246
	CouponGiftStopResult(request *CouponGiftStopRequest) ([]byte, error)
	// 礼金停止
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftGetResult(request *CouponGiftGetRequest) ([]byte, error)
	// 礼金停止
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftStopResultByKey(giftCouponKey string) ([]byte, error)
	// 礼金效果数据
	//    文档: https://union.jd.com/openplatform/api/12248
	StatisticGiftCouponQueryResult(request *StatisticGiftCouponQueryRequest) ([]byte, error)

	// 礼金创建
	//    文档: https://union.jd.com/openplatform/api/12246
	CouponGiftStopMap(request *CouponGiftStopRequest) (map[string]interface{}, error)
	// 礼金停止
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftGetMap(request *CouponGiftGetRequest) (map[string]interface{}, error)
	// 礼金停止
	//    文档: https://union.jd.com/openplatform/api/12240
	CouponGiftStopMapByKey(giftCouponKey string) (map[string]interface{}, error)
	// 礼金效果数据
	//    文档: https://union.jd.com/openplatform/api/12248
	StatisticGiftCouponQueryMap(request *StatisticGiftCouponQueryRequest) (map[string]interface{}, error)
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
	err := p.service.Do(&res, CouponGiftGet, param)
	return &res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStop(request *CouponGiftStopRequest) (*CouponGiftStopResult, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftStopResult
	err := p.service.Do(&res, CouponGiftStop, param)
	return &res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopByKey(giftCouponKey string) (*CouponGiftStopResult, error) {
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
	err := p.service.Do(&res, StatisticGiftCouponQuery, param)
	return &res, err
}

// 礼金停止
// Deprecated: 使用 CouponGiftStopByKey
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopBy(giftCouponKey string) (*CouponGiftStopResult, error) {
	return p.CouponGiftStop(&CouponGiftStopRequest{
		GiftCouponKey: giftCouponKey,
	})
}

// 礼金创建
//    文档: https://union.jd.com/openplatform/api/12246
func (p *GiftServiceImpl) CouponGiftGetResult(request *CouponGiftGetRequest) ([]byte, error) {
	res, err := p.CouponGiftGet(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopResult(request *CouponGiftStopRequest) ([]byte, error) {
	res, err := p.CouponGiftStop(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopResultByKey(giftCouponKey string) ([]byte, error) {
	res, err := p.CouponGiftStopByKey(giftCouponKey)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 礼金效果数据
//    文档: https://union.jd.com/openplatform/api/12248
func (p *GiftServiceImpl) StatisticGiftCouponQueryResult(request *StatisticGiftCouponQueryRequest) ([]byte, error) {
	res, err := p.StatisticGiftCouponQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 礼金创建
//    文档: https://union.jd.com/openplatform/api/12246
func (p *GiftServiceImpl) CouponGiftGetMap(request *CouponGiftGetRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, CouponGiftGet, param)
	return res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopMap(request *CouponGiftStopRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, CouponGiftStop, param)
	return res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (p *GiftServiceImpl) CouponGiftStopMapByKey(giftCouponKey string) (map[string]interface{}, error) {
	return p.CouponGiftStopMap(&CouponGiftStopRequest{
		GiftCouponKey: giftCouponKey,
	})
}

// 礼金效果数据
//    文档: https://union.jd.com/openplatform/api/12248
func (p *GiftServiceImpl) StatisticGiftCouponQueryMap(request *StatisticGiftCouponQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, StatisticGiftCouponQuery, param)
	return res, err
}
