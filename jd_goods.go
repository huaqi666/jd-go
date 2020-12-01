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
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetQuery(*CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error)
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
	//    文档: https://union.jd.com/openplatform/api/10434
	CategoryGoodsGetQueryResult(*CategoryGoodsGetRequest) ([]byte, error)
	// 商品详情查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11248
	GoodsGigfieldQueryResult(request *GoodsGigFieldQueryRequest) ([]byte, error)
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
func (g *GoodsServiceImpl) GoodsJingfenQuery(request *GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error) {
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsJingfenQueryResult
	err := g.service.Do(&res, GoodsJingfenQuery, param)
	return &res, err
}

// 关键词商品查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10420
func (g *GoodsServiceImpl) GoodsQuery(request *GoodsQueryRequest) (*GoodsQueryResult, error) {
	param := map[string]interface{}{}
	param["goodsReqDTO"] = request
	var res GoodsQueryResult
	err := g.service.Do(&res, GoodsQuery, param)
	return &res, err
}

/* 根据skuid查询商品信息接口
   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
//    文档: https://union.jd.com/openplatform/api/10422
func (g *GoodsServiceImpl) GoodsPromotiongoodsinfoQuery(skuIds string) (*GoodsPromotiongoodsinfoQueryResult, error) {
	param := map[string]interface{}{}
	param["skuIds"] = skuIds
	var res GoodsPromotiongoodsinfoQueryResult
	err := g.service.Do(&res, GoodsPromotiongoodsinfoQuery, param)
	return &res, err
}

// 商品类目查询接口
//    文档: https://union.jd.com/openplatform/api/10434
func (g *GoodsServiceImpl) CategoryGoodsGetQuery(request *CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error) {
	param := map[string]interface{}{}
	param["req"] = request
	var res CategoryGoodsGetResult
	err := g.service.Do(&res, CategoryGoodsGetQuery, param)
	return &res, err
}

// 商品详情查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11248
func (g *GoodsServiceImpl) GoodsGigfieldQuery(request *GoodsGigFieldQueryRequest) (*GoodsGigFieldQueryResult, error) {
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsGigFieldQueryResult
	err := g.service.Do(&res, GoodsGigfieldQuery, param)
	return &res, err
}

// 京粉精选商品查询接口
//    文档: https://union.jd.com/openplatform/api/10417
func (g *GoodsServiceImpl) GoodsJingfenQueryResult(request *GoodsJingfenQueryRequest) ([]byte, error) {
	res, err := g.GoodsJingfenQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 关键词商品查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/10420
func (g *GoodsServiceImpl) GoodsQueryResult(request *GoodsQueryRequest) ([]byte, error) {
	res, err := g.GoodsQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

/* 根据skuid查询商品信息接口
   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息）*/
//    文档: https://union.jd.com/openplatform/api/10422
func (g *GoodsServiceImpl) GoodsPromotiongoodsinfoQueryResult(skuIds string) ([]byte, error) {
	res, err := g.GoodsPromotiongoodsinfoQuery(skuIds)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 商品类目查询接口
//    文档: https://union.jd.com/openplatform/api/10434
func (g *GoodsServiceImpl) CategoryGoodsGetQueryResult(request *CategoryGoodsGetRequest) ([]byte, error) {
	res, err := g.CategoryGoodsGetQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 商品详情查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11248
func (g *GoodsServiceImpl) GoodsGigfieldQueryResult(request *GoodsGigFieldQueryRequest) ([]byte, error) {
	res, err := g.GoodsGigfieldQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}
