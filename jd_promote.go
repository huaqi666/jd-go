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
	// 通过小程序获取推广链接【申请】
	//    文档: https://union.jd.com/openplatform/api/11793
	PromotionAppletGet(request *PromotionAppletGetRequest) (*PromotionAppletGetResult, error)

	// 网站/APP获取推广链接接口
	//    文档: https://union.jd.com/openplatform/api/10421
	PromotionCommonGetResult(*PromoteCommonGetRequest) ([]byte, error)
	// 社交媒体获取推广链接接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10424
	PromotionBysubunionidGetResult(*PromotionBysubunionidGetRequest) ([]byte, error)
	// 工具商获取推广链接接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10425
	PromotionByunionidGetResult(*PromotionByunionidGetRequest) ([]byte, error)
	// 通过小程序获取推广链接【申请】
	//    文档: https://union.jd.com/openplatform/api/11793
	PromotionAppletGetResult(request *PromotionAppletGetRequest) ([]byte, error)

	// 网站/APP获取推广链接接口
	//    文档: https://union.jd.com/openplatform/api/10421
	PromotionCommonGetMap(*PromoteCommonGetRequest) (map[string]interface{}, error)
	// 社交媒体获取推广链接接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10424
	PromotionBysubunionidGetMap(*PromotionBysubunionidGetRequest) (map[string]interface{}, error)
	// 工具商获取推广链接接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10425
	PromotionByunionidGetMap(*PromotionByunionidGetRequest) (map[string]interface{}, error)
	// 通过小程序获取推广链接【申请】
	//    文档: https://union.jd.com/openplatform/api/11793
	PromotionAppletGetMap(request *PromotionAppletGetRequest) (map[string]interface{}, error)
}

type PromotionServiceImpl struct {
	service Service
}

func newPromotionService(service Service) PromotionService {
	return &PromotionServiceImpl{
		service: service,
	}
}

// 网站/APP获取推广链接接口
//    文档: https://union.jd.com/openplatform/api/10421
func (promo *PromotionServiceImpl) PromotionCommonGet(request *PromoteCommonGetRequest) (*PromoteCommonGetResult, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromoteCommonGetResult
	err = promo.service.Do(&res, PromotionCommonGet, param)
	return &res, err
}

// 社交媒体获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10424
func (promo *PromotionServiceImpl) PromotionBysubunionidGet(request *PromotionBysubunionidGetRequest) (*PromotionBysubunionidGetResult, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionBysubunionidGetResult
	err = promo.service.Do(&res, PromotionBysubunionidGet, param)
	return &res, err
}

// 工具商获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10425
func (promo *PromotionServiceImpl) PromotionByunionidGet(request *PromotionByunionidGetRequest) (*PromotionByunionidGetResult, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionByunionidGetResult
	err = promo.service.Do(&res, PromotionByunionidGet, param)
	return &res, err
}

// 通过小程序获取推广链接【申请】
//    文档: https://union.jd.com/openplatform/api/11793
func (promo *PromotionServiceImpl) PromotionAppletGet(request *PromotionAppletGetRequest) (*PromotionAppletGetResult, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionAppletGetResult
	err = promo.service.Do(&res, PromotionAppletGet, param)
	return &res, err
}

// 网站/APP获取推广链接接口
//    文档: https://union.jd.com/openplatform/api/10421
func (promo *PromotionServiceImpl) PromotionCommonGetResult(request *PromoteCommonGetRequest) ([]byte, error) {
	res, err := promo.PromotionCommonGet(request)
	return promo.service.GetResult(res, err)
}

// 社交媒体获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10424
func (promo *PromotionServiceImpl) PromotionBysubunionidGetResult(request *PromotionBysubunionidGetRequest) ([]byte, error) {
	res, err := promo.PromotionBysubunionidGet(request)
	return promo.service.GetResult(res, err)
}

// 工具商获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10425
func (promo *PromotionServiceImpl) PromotionByunionidGetResult(request *PromotionByunionidGetRequest) ([]byte, error) {
	res, err := promo.PromotionByunionidGet(request)
	return promo.service.GetResult(res, err)
}

// 通过小程序获取推广链接【申请】
//    文档: https://union.jd.com/openplatform/api/11793
func (promo *PromotionServiceImpl) PromotionAppletGetResult(request *PromotionAppletGetRequest) ([]byte, error) {
	res, err := promo.PromotionAppletGet(request)
	return promo.service.GetResult(res, err)
}

// 网站/APP获取推广链接接口
//    文档: https://union.jd.com/openplatform/api/10421
func (promo *PromotionServiceImpl) PromotionCommonGetMap(request *PromoteCommonGetRequest) (map[string]interface{}, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res = make(map[string]interface{})
	err = promo.service.Do(&res, PromotionCommonGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if promo.service.GetRouteApi() == JosRootEndpoint {
		key = PromotionCommonGetResponce
	} else {
		key = PromotionCommonGetResponse
	}
	return promo.service.ParseResult(res, key)
}

// 社交媒体获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10424
func (promo *PromotionServiceImpl) PromotionBysubunionidGetMap(request *PromotionBysubunionidGetRequest) (map[string]interface{}, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res map[string]interface{}
	err = promo.service.Do(&res, PromotionBysubunionidGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if promo.service.GetRouteApi() == JosRootEndpoint {
		key = PromotionBysubunionidGetResponce
	} else {
		key = PromotionBysubunionidGetResponse
	}
	return promo.service.ParseResult(res, key)
}

// 工具商获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10425
func (promo *PromotionServiceImpl) PromotionByunionidGetMap(request *PromotionByunionidGetRequest) (map[string]interface{}, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res map[string]interface{}
	err = promo.service.Do(&res, PromotionByunionidGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if promo.service.GetRouteApi() == JosRootEndpoint {
		key = PromotionByunionidGetResponce
	} else {
		key = PromotionByunionidGetResponse
	}
	return promo.service.ParseResult(res, key)
}

// 通过小程序获取推广链接【申请】
//    文档: https://union.jd.com/openplatform/api/11793
func (promo *PromotionServiceImpl) PromotionAppletGetMap(request *PromotionAppletGetRequest) (map[string]interface{}, error) {
	err := promo.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res map[string]interface{}
	err = promo.service.Do(&res, PromotionAppletGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if promo.service.GetRouteApi() == JosRootEndpoint {
		key = PromotionAppletGetResponce
	} else {
		key = PromotionAppletGetResponse
	}
	return promo.service.ParseResult(res, key)
}
