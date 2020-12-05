// jd联盟go-sdk
package jd

type GoodsService interface {
	// 京粉精选商品查询接口
	//    文档: https://union.jd.com/openplatform/api/10417
	GoodsJingfenQuery(*GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error)
	// 关键词商品查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10420
	GoodsQuery(*GoodsQueryRequest) (*GoodsQueryResult, error)
	/* 根据skuid查询商品信息接口
	   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
	  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
	//    文档: https://union.jd.com/openplatform/api/10422
	GoodsPromotiongoodsinfoQuery(skuIds string) (*GoodsPromotiongoodsinfoQueryResult, error)
	// 商品类目查询接口
	// Deprecated: 使用新接口: CategoryGoodsGet
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetQuery(*CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error)
	// 商品类目查询接口
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGet(*CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error)
	// 商品详情查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11248
	GoodsGigfieldQuery(*GoodsGigFieldQueryRequest) (*GoodsGigFieldQueryResult, error)

	// 京粉精选商品查询接口
	//    文档: https://union.jd.com/openplatform/api/10417
	GoodsJingfenQueryResult(*GoodsJingfenQueryRequest) ([]byte, error)
	// 关键词商品查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10420
	GoodsQueryResult(*GoodsQueryRequest) ([]byte, error)
	/* 根据skuid查询商品信息接口
	   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
	  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
	//    文档: https://union.jd.com/openplatform/api/10422
	GoodsPromotiongoodsinfoQueryResult(skuIds string) ([]byte, error)
	// 商品类目查询接口
	// Deprecated: 使用新接口: GoodsService.CategoryGoodsGetResult
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetQueryResult(*CategoryGoodsGetRequest) ([]byte, error)
	// 商品类目查询接口
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetResult(*CategoryGoodsGetRequest) ([]byte, error)
	// 商品详情查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11248
	GoodsGigfieldQueryResult(request *GoodsGigFieldQueryRequest) ([]byte, error)

	// 京粉精选商品查询接口
	//    文档: https://union.jd.com/openplatform/api/10417
	GoodsJingfenQueryMap(*GoodsJingfenQueryRequest) (map[string]interface{}, error)
	// 关键词商品查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/10420
	GoodsQueryMap(*GoodsQueryRequest) (map[string]interface{}, error)
	/* 根据skuid查询商品信息接口
	   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
	  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
	//    文档: https://union.jd.com/openplatform/api/10422
	GoodsPromotiongoodsinfoQueryMap(skuIds string) (map[string]interface{}, error)
	// 商品类目查询接口
	// Deprecated: 使用新接口: CategoryGoodsGetMap
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetQueryMap(*CategoryGoodsGetRequest) (map[string]interface{}, error)
	// 商品类目查询接口
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetMap(*CategoryGoodsGetRequest) (map[string]interface{}, error)
	// 商品详情查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11248
	GoodsGigfieldQueryMap(request *GoodsGigFieldQueryRequest) (map[string]interface{}, error)
}

type GoodsServiceImpl struct {
	service Service
}

func newGoodsService(service Service) GoodsService {
	return &GoodsServiceImpl{
		service: service,
	}
}

// 京粉精选商品查询接口
//    文档: https://union.jd.com/openplatform/api/10417
func (goods *GoodsServiceImpl) GoodsJingfenQuery(request *GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsJingfenQueryResult
	err = goods.service.Do(&res, GoodsJingfenQuery, param)
	return &res, err
}

// 关键词商品查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10420
func (goods *GoodsServiceImpl) GoodsQuery(request *GoodsQueryRequest) (*GoodsQueryResult, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["goodsReqDTO"] = request
	var res GoodsQueryResult
	err = goods.service.Do(&res, GoodsQuery, param)
	return &res, err
}

/* 根据skuid查询商品信息接口
   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
//    文档: https://union.jd.com/openplatform/api/10422
func (goods *GoodsServiceImpl) GoodsPromotiongoodsinfoQuery(skuIds string) (*GoodsPromotiongoodsinfoQueryResult, error) {
	err := goods.service.CheckRequiredParameters(struct{ skuIds string }{skuIds: skuIds})
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["skuIds"] = skuIds
	var res GoodsPromotiongoodsinfoQueryResult
	err = goods.service.Do(&res, GoodsPromotiongoodsinfoQuery, param)
	return &res, err
}

// 商品类目查询接口
// Deprecated: 使用新接口: CategoryGoodsGet
//    文档: https://union.jd.com/openplatform/api/10434
func (goods *GoodsServiceImpl) CategoryGoodsGetQuery(request *CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error) {
	return goods.CategoryGoodsGet(request)
}

// 商品类目查询接口
//    文档: https://union.jd.com/openplatform/api/10434
func (goods *GoodsServiceImpl) CategoryGoodsGet(request *CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["req"] = request
	var res CategoryGoodsGetResult
	err = goods.service.Do(&res, CategoryGoodsGet, param)
	return &res, err
}

// 商品详情查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11248
func (goods *GoodsServiceImpl) GoodsGigfieldQuery(request *GoodsGigFieldQueryRequest) (*GoodsGigFieldQueryResult, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsGigFieldQueryResult
	err = goods.service.Do(&res, GoodsGigfieldQuery, param)
	return &res, err
}

// 京粉精选商品查询接口
//    文档: https://union.jd.com/openplatform/api/10417
func (goods *GoodsServiceImpl) GoodsJingfenQueryResult(request *GoodsJingfenQueryRequest) ([]byte, error) {
	res, err := goods.GoodsJingfenQuery(request)
	return goods.service.GetResult(res, err)
}

// 关键词商品查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10420
func (goods *GoodsServiceImpl) GoodsQueryResult(request *GoodsQueryRequest) ([]byte, error) {
	res, err := goods.GoodsQuery(request)
	return goods.service.GetResult(res, err)
}

/* 根据skuid查询商品信息接口
   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
//    文档: https://union.jd.com/openplatform/api/10422
func (goods *GoodsServiceImpl) GoodsPromotiongoodsinfoQueryResult(skuIds string) ([]byte, error) {
	res, err := goods.GoodsPromotiongoodsinfoQuery(skuIds)
	return goods.service.GetResult(res, err)
}

// 商品类目查询接口
// Deprecated: 使用新接口: GoodsService.CategoryGoodsGetResult
//    文档: https://union.jd.com/openplatform/api/10434
func (goods *GoodsServiceImpl) CategoryGoodsGetQueryResult(request *CategoryGoodsGetRequest) ([]byte, error) {
	return goods.CategoryGoodsGetResult(request)
}

// 商品类目查询接口
//    文档: https://union.jd.com/openplatform/api/10434
func (goods *GoodsServiceImpl) CategoryGoodsGetResult(request *CategoryGoodsGetRequest) ([]byte, error) {
	res, err := goods.CategoryGoodsGet(request)
	return goods.service.GetResult(res, err)
}

// 商品详情查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11248
func (goods *GoodsServiceImpl) GoodsGigfieldQueryResult(request *GoodsGigFieldQueryRequest) ([]byte, error) {
	res, err := goods.GoodsGigfieldQuery(request)
	return goods.service.GetResult(res, err)
}

// 京粉精选商品查询接口
//    文档: https://union.jd.com/openplatform/api/10417
func (goods *GoodsServiceImpl) GoodsJingfenQueryMap(request *GoodsJingfenQueryRequest) (map[string]interface{}, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res map[string]interface{}
	err = goods.service.Do(&res, GoodsJingfenQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if goods.service.GetRouteApi() == JosRootEndpoint {
		key = GoodsJingfenQueryResponce
	} else {
		key = GoodsJingfenQueryResponse
	}
	return goods.service.ParseResult(res, key)
}

// 关键词商品查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10420
func (goods *GoodsServiceImpl) GoodsQueryMap(request *GoodsQueryRequest) (map[string]interface{}, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["goodsReqDTO"] = request
	var res map[string]interface{}
	err = goods.service.Do(&res, GoodsQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if goods.service.GetRouteApi() == JosRootEndpoint {
		key = GoodsQueryResponce
	} else {
		key = GoodsQueryResponse
	}
	return goods.service.ParseResult(res, key)
}

/* 根据skuid查询商品信息接口
   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
//    文档: https://union.jd.com/openplatform/api/10422
func (goods *GoodsServiceImpl) GoodsPromotiongoodsinfoQueryMap(skuIds string) (map[string]interface{}, error) {
	err := goods.service.CheckRequiredParameters(struct{ skuIds string }{skuIds: skuIds})
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["skuIds"] = skuIds
	var res map[string]interface{}
	err = goods.service.Do(&res, GoodsPromotiongoodsinfoQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if goods.service.GetRouteApi() == JosRootEndpoint {
		key = GoodsPromotiongoodsinfoQueryResponce
	} else {
		key = GoodsPromotiongoodsinfoQueryResponse
	}
	return goods.service.ParseResult(res, key)
}

// 商品类目查询接口
// Deprecated: 使用新接口: CategoryGoodsGetMap
//    文档: https://union.jd.com/openplatform/api/10434
func (goods *GoodsServiceImpl) CategoryGoodsGetQueryMap(request *CategoryGoodsGetRequest) (map[string]interface{}, error) {
	return goods.CategoryGoodsGetMap(request)
}

// 商品类目查询接口
//    文档: https://union.jd.com/openplatform/api/10434
func (goods *GoodsServiceImpl) CategoryGoodsGetMap(request *CategoryGoodsGetRequest) (map[string]interface{}, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["req"] = request
	var res map[string]interface{}
	err = goods.service.Do(&res, CategoryGoodsGet, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if goods.service.GetRouteApi() == JosRootEndpoint {
		key = CategoryGoodsGetResponce
	} else {
		key = CategoryGoodsGetResponse
	}
	return goods.service.ParseResult(res, key)
}

// 商品详情查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11248
func (goods *GoodsServiceImpl) GoodsGigfieldQueryMap(request *GoodsGigFieldQueryRequest) (map[string]interface{}, error) {
	err := goods.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res map[string]interface{}
	err = goods.service.Do(&res, GoodsGigfieldQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if goods.service.GetRouteApi() == JosRootEndpoint {
		key = GoodsGigfieldQueryResponce
	} else {
		key = GoodsGigfieldQueryResponse
	}
	return goods.service.ParseResult(res, key)
}
