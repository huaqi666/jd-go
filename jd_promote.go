// jd联盟go-sdk
package jd

import "encoding/json"

type PromoteService interface {
	// 网站/APP获取推广链接接口
	CommonGet(*PromoteCommonGetRequest) (*PromoteCommonGetResult, error)
}

type PromoteServiceImpl struct {
	service Service
}

func newPromoteService(service Service) PromoteService {
	return &PromoteServiceImpl{
		service: service,
	}
}

func (p *PromoteServiceImpl) CommonGet(request *PromoteCommonGetRequest) (*PromoteCommonGetResult, error) {
	param := map[string]string{}
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	param["promotionCodeReq"] = string(b)
	var res PromoteCommonGetResult
	err = p.service.DoGet(&res, PromotionCommonGet, param)
	return &res, err
}
