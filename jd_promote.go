// jd联盟go-sdk
package jd

type PromotionService interface {
	// 网站/APP获取推广链接接口
	//    文档: https://union.jd.com/openplatform/api/10421
	PromotionCommonGet(*PromoteCommonGetRequest) (*PromoteCommonGetResult, error)
	// 社交媒体获取推广链接接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10424
	PromotionBysubunionidGet(*PromotionBysubunionidGetRequest) (*PromotionBysubunionidGetResult, error)
	// 工具商获取推广链接接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10425
	PromotionByunionidGet(*PromotionByunionidGetRequest) (*PromotionByunionidGetResult, error)
}

type PromotionServiceImpl struct {
	service Service
}

func newPromotionService(service Service) PromotionService {
	return &PromotionServiceImpl{
		service: service,
	}
}

func (p *PromotionServiceImpl) PromotionCommonGet(request *PromoteCommonGetRequest) (*PromoteCommonGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromoteCommonGetResult
	err := p.service.Do(&res, PromotionCommonGet, param)
	return &res, err
}

func (p *PromotionServiceImpl) PromotionBysubunionidGet(request *PromotionBysubunionidGetRequest) (*PromotionBysubunionidGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionBysubunionidGetResult
	err := p.service.Do(&res, PromotionBysubunionidQuery, param)
	return &res, err
}

func (p *PromotionServiceImpl) PromotionByunionidGet(request *PromotionByunionidGetRequest) (*PromotionByunionidGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionByunionidGetResult
	err := p.service.Do(&res, PromotionByunionidQuery, param)
	return &res, err
}
