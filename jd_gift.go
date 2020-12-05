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
func (gift *GiftServiceImpl) CouponGiftGet(request *CouponGiftGetRequest) (*CouponGiftGetResult, error) {
	err := gift.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftGetResult
	err = gift.service.Do(&res, CouponGiftGet, param)
	return &res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStop(request *CouponGiftStopRequest) (*CouponGiftStopResult, error) {
	err := gift.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res CouponGiftStopResult
	err = gift.service.Do(&res, CouponGiftStop, param)
	return &res, err
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStopByKey(giftCouponKey string) (*CouponGiftStopResult, error) {
	return gift.CouponGiftStop(&CouponGiftStopRequest{
		GiftCouponKey: giftCouponKey,
	})
}

// 礼金效果数据
//    文档: https://union.jd.com/openplatform/api/12248
func (gift *GiftServiceImpl) StatisticGiftCouponQuery(request *StatisticGiftCouponQueryRequest) (*StatisticGiftCouponQueryResult, error) {
	err := gift.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res StatisticGiftCouponQueryResult
	err = gift.service.Do(&res, StatisticGiftCouponQuery, param)
	return &res, err
}

// 礼金停止
// Deprecated: 使用 CouponGiftStopByKey
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStopBy(giftCouponKey string) (*CouponGiftStopResult, error) {
	return gift.CouponGiftStop(&CouponGiftStopRequest{
		GiftCouponKey: giftCouponKey,
	})
}

// 礼金创建
//    文档: https://union.jd.com/openplatform/api/12246
func (gift *GiftServiceImpl) CouponGiftGetResult(request *CouponGiftGetRequest) ([]byte, error) {
	res, err := gift.CouponGiftGet(request)
	return gift.service.GetResult(res, err)
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStopResult(request *CouponGiftStopRequest) ([]byte, error) {
	res, err := gift.CouponGiftStop(request)
	return gift.service.GetResult(res, err)
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStopResultByKey(giftCouponKey string) ([]byte, error) {
	res, err := gift.CouponGiftStopByKey(giftCouponKey)
	return gift.service.GetResult(res, err)
}

// 礼金效果数据
//    文档: https://union.jd.com/openplatform/api/12248
func (gift *GiftServiceImpl) StatisticGiftCouponQueryResult(request *StatisticGiftCouponQueryRequest) ([]byte, error) {
	res, err := gift.StatisticGiftCouponQuery(request)
	return gift.service.GetResult(res, err)
}

// 礼金创建
//    文档: https://union.jd.com/openplatform/api/12246
func (gift *GiftServiceImpl) CouponGiftGetMap(request *CouponGiftGetRequest) (map[string]interface{}, error) {
	err := gift.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res map[string]interface{}
	err = gift.service.Do(&res, CouponGiftGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if gift.service.GetRouteApi() == JosRootEndpoint {
		key = CouponGiftGetResponce
	} else {
		key = CouponGiftGetResponse
	}
	return gift.service.ParseResult(res, key)
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStopMap(request *CouponGiftStopRequest) (map[string]interface{}, error) {
	err := gift.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["couponReq"] = request
	var res map[string]interface{}
	err = gift.service.Do(&res, CouponGiftStop, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if gift.service.GetRouteApi() == JosRootEndpoint {
		key = CouponGiftStopResponce
	} else {
		key = CouponGiftStopResponse
	}
	return gift.service.ParseResult(res, key)
}

// 礼金停止
//    文档: https://union.jd.com/openplatform/api/12240
func (gift *GiftServiceImpl) CouponGiftStopMapByKey(giftCouponKey string) (map[string]interface{}, error) {
	return gift.CouponGiftStopMap(&CouponGiftStopRequest{
		GiftCouponKey: giftCouponKey,
	})
}

// 礼金效果数据
//    文档: https://union.jd.com/openplatform/api/12248
func (gift *GiftServiceImpl) StatisticGiftCouponQueryMap(request *StatisticGiftCouponQueryRequest) (map[string]interface{}, error) {
	err := gift.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["effectDataReq"] = request
	var res map[string]interface{}
	err = gift.service.Do(&res, StatisticGiftCouponQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if gift.service.GetRouteApi() == JosRootEndpoint {
		key = StatisticGiftCouponQueryResponce
	} else {
		key = StatisticGiftCouponQueryResponse
	}
	return gift.service.ParseResult(res, key)
}
