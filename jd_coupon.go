// jd联盟go-sdk
package jd

type CouponService interface {
	// 优惠券领取情况查询接口【申请】
	CouponQuery(*CouponQueryRequest) (*CouponQueryResult, error)
}

type CouponServiceImpl struct {
	service Service
}

func newCouponService(service Service) CouponService {

	return &CouponServiceImpl{
		service: service,
	}
}

func (c CouponServiceImpl) CouponQuery(request *CouponQueryRequest) (*CouponQueryResult, error) {
	param := map[string]interface{}{}
	param["couponUrls"] = request
	var res CouponQueryResult
	err := c.service.Do(&res, CouponQuery, param)
	return &res, err
}
