// jd联盟go-sdk
package jd

type CouponService interface {
	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQuery([]*CouponQueryRequest) (*CouponQueryResult, error)
	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryByUrls(urls ...string) (*CouponQueryResult, error)

	// 优惠券领取情况查询接口【申请】
	// Deprecated: 使用 CouponQueryByUrls
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryByList(urls ...string) (*CouponQueryResult, error)

	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryResult(request []*CouponQueryRequest) ([]byte, error)
	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryResultByUrls(urls ...string) ([]byte, error)

	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryMap(request []*CouponQueryRequest) (map[string]interface{}, error)
	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryMapByUrls(urls ...string) (map[string]interface{}, error)
}

type CouponServiceImpl struct {
	service Service
}

func newCouponService(service Service) CouponService {

	return &CouponServiceImpl{
		service: service,
	}
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQuery(request []*CouponQueryRequest) (*CouponQueryResult, error) {
	err := cou.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["couponUrls"] = request
	var res CouponQueryResult
	err = cou.service.Do(&res, CouponQuery, param)
	return &res, err
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQueryByUrls(urls ...string) (*CouponQueryResult, error) {
	var arr []*CouponQueryRequest
	for _, url := range urls {
		arr = append(arr, &CouponQueryRequest{
			CouponUrl: url,
		})
	}
	return cou.CouponQuery(arr)
}

// 优惠券领取情况查询接口【申请】
// Deprecated: 使用 CouponQueryByUrls
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQueryByList(urls ...string) (*CouponQueryResult, error) {
	return cou.CouponQueryByUrls(urls...)
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQueryResult(request []*CouponQueryRequest) ([]byte, error) {
	res, err := cou.CouponQuery(request)
	return cou.service.GetResult(res, err)
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQueryResultByUrls(urls ...string) ([]byte, error) {
	return cou.service.GetResult(cou.CouponQueryByUrls(urls...))
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQueryMap(request []*CouponQueryRequest) (map[string]interface{}, error) {
	err := cou.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["couponUrls"] = request
	var res map[string]interface{}
	err = cou.service.Do(&res, CouponQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if cou.service.GetRouteApi() == JosRootEndpoint {
		key = CouponQueryResponce
	} else {
		key = CouponQueryResponse
	}
	return cou.service.ParseResult(res, key)
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (cou *CouponServiceImpl) CouponQueryMapByUrls(urls ...string) (map[string]interface{}, error) {
	var arr []*CouponQueryRequest
	for _, url := range urls {
		arr = append(arr, &CouponQueryRequest{
			CouponUrl: url,
		})
	}
	return cou.CouponQueryMap(arr)
}
