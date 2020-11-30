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
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryResult(request []*CouponQueryRequest) ([]byte, error)
	// 优惠券领取情况查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryResultByUrls(urls ...string) ([]byte, error)
	// 优惠券领取情况查询接口【申请】
	// Deprecated: 使用 CouponQueryByUrls
	//    文档: https://union.jd.com/openplatform/api/10423
	CouponQueryByList(urls ...string) (*CouponQueryResult, error)
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
// Deprecated: 使用 CouponQueryByUrls
//    文档: https://union.jd.com/openplatform/api/10423
func (c *CouponServiceImpl) CouponQueryByList(urls ...string) (*CouponQueryResult, error) {
	return c.CouponQueryByUrls(urls...)
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (c *CouponServiceImpl) CouponQueryByUrls(urls ...string) (*CouponQueryResult, error) {
	var arr []*CouponQueryRequest
	for _, url := range urls {
		arr = append(arr, &CouponQueryRequest{
			CouponUrl: url,
		})
	}
	return c.CouponQuery(arr)
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (c *CouponServiceImpl) CouponQuery(request []*CouponQueryRequest) (*CouponQueryResult, error) {
	param := map[string]interface{}{}
	param["couponUrls"] = request
	var res CouponQueryResult
	err := c.service.Do(&res, CouponQuery, param)
	return &res, err
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (c *CouponServiceImpl) CouponQueryResultByUrls(urls ...string) ([]byte, error) {
	res, err := c.CouponQueryByUrls(urls...)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 优惠券领取情况查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10423
func (c *CouponServiceImpl) CouponQueryResult(request []*CouponQueryRequest) ([]byte, error) {
	res, err := c.CouponQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}
