// jd联盟go-sdk
package jd

import "encoding/json"

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
	param := map[string]string{}
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	param["couponUrls"] = string(b)
	var res CouponQueryResult
	err = c.service.DoGet(&res, CouponQuery, param)
	return &res, err
}
