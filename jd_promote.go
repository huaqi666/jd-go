// jd联盟go-sdk
package jd

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
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromoteCommonGetResult
	err := p.service.Do(&res, PromotionCommonGet, param)
	return &res, err
}
