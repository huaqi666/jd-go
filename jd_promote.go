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
func (p *PromotionServiceImpl) PromotionCommonGet(request *PromoteCommonGetRequest) (*PromoteCommonGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromoteCommonGetResult
	err := p.service.Do(&res, PromotionCommonGet, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 社交媒体获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10424
func (p *PromotionServiceImpl) PromotionBysubunionidGet(request *PromotionBysubunionidGetRequest) (*PromotionBysubunionidGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionBysubunionidGetResult
	err := p.service.Do(&res, PromotionBysubunionidQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 工具商获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10425
func (p *PromotionServiceImpl) PromotionByunionidGet(request *PromotionByunionidGetRequest) (*PromotionByunionidGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionByunionidGetResult
	err := p.service.Do(&res, PromotionByunionidQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 通过小程序获取推广链接【申请】
//    文档: https://union.jd.com/openplatform/api/11793
func (p *PromotionServiceImpl) PromotionAppletGet(request *PromotionAppletGetRequest) (*PromotionAppletGetResult, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res PromotionAppletGetResult
	err := p.service.Do(&res, PromotionAppletQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 网站/APP获取推广链接接口
//    文档: https://union.jd.com/openplatform/api/10421
func (p *PromotionServiceImpl) PromotionCommonGetResult(request *PromoteCommonGetRequest) ([]byte, error) {
	res, err := p.PromotionCommonGet(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 社交媒体获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10424
func (p *PromotionServiceImpl) PromotionBysubunionidGetResult(request *PromotionBysubunionidGetRequest) ([]byte, error) {
	res, err := p.PromotionBysubunionidGet(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 工具商获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10425
func (p *PromotionServiceImpl) PromotionByunionidGetResult(request *PromotionByunionidGetRequest) ([]byte, error) {
	res, err := p.PromotionByunionidGet(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 通过小程序获取推广链接【申请】
//    文档: https://union.jd.com/openplatform/api/11793
func (p *PromotionServiceImpl) PromotionAppletGetResult(request *PromotionAppletGetRequest) ([]byte, error) {
	res, err := p.PromotionAppletGet(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 网站/APP获取推广链接接口
//    文档: https://union.jd.com/openplatform/api/10421
func (p *PromotionServiceImpl) PromotionCommonGetMap(request *PromoteCommonGetRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res = make(map[string]interface{})
	err := p.service.Do(&res, PromotionCommonGet, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}

// 社交媒体获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10424
func (p *PromotionServiceImpl) PromotionBysubunionidGetMap(request *PromotionBysubunionidGetRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, PromotionBysubunionidQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}

// 工具商获取推广链接接口【申请】
//    文档: https://union.jd.com/openplatform/api/10425
func (p *PromotionServiceImpl) PromotionByunionidGetMap(request *PromotionByunionidGetRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, PromotionByunionidQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}

// 通过小程序获取推广链接【申请】
//    文档: https://union.jd.com/openplatform/api/11793
func (p *PromotionServiceImpl) PromotionAppletGetMap(request *PromotionAppletGetRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["promotionCodeReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, PromotionAppletQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}
