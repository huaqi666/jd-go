// jd联盟go-sdk
package jd

type PromotionService interface {
	// 网站/APP获取推广链接接口
	PromotionCommonGet(*PromoteCommonGetRequest) (*PromoteCommonGetResult, error)
	// 社交媒体获取推广链接接口【申请】
	PromotionBysubunionidQuery(*PromotionBysubunionidQueryRequest) (*PromotionBysubunionidQueryResult, error)
	// 工具商获取推广链接接口【申请】
	PromotionByunionidQuery(*PromotionByunionidQueryRequest) (*PromotionByunionidQueryResult, error)
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

func (p *PromotionServiceImpl) PromotionBysubunionidQuery(request *PromotionBysubunionidQueryRequest) (*PromotionBysubunionidQueryResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionBysubunionidQueryResult
	err := p.service.Do(&res, PromotionBysubunionidQuery, param)
	return &res, err
}

func (p *PromotionServiceImpl) PromotionByunionidQuery(request *PromotionByunionidQueryRequest) (*PromotionByunionidQueryResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionByunionidQueryResult
	err := p.service.Do(&res, PromotionByunionidQuery, param)
	return &res, err
}
