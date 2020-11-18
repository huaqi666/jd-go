// jd联盟go-sdk
package jd

type GoodsService interface {
	// 京粉精选商品查询接口
	GoodsJingfenQuery(*GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error)
	// 关键词商品查询接口【申请】
	GoodsQuery(*GoodsQueryRequest) (*GoodsQueryResult, error)
	/* 根据skuid查询商品信息接口
	   skuIds: 京东skuID串，逗号分割，最多100个，开发示例如param_json={'skuIds':'5225346,7275691'}
	  （非常重要 请大家关注：如果输入的sk串中某个skuID的商品不在推广中[就是没有佣金]，返回结果中不会包含这个商品的信息） */
	GoodsPromotiongoodsinfoQuery(skuIds string) (*GoodsPromotiongoodsinfoResult, error)
	// 商品类目查询接口
	CategoryGoodsGetQuery(*CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error)
	// 商品详情查询接口【申请】
	GoodsGigfieldQuery(*GoodsGigfieldRequest) (*GoodsGigfieldResult, error)
}

type GoodsServiceImpl struct {
	service Service
}

func newGoodsService(service Service) GoodsService {
	return &GoodsServiceImpl{
		service: service,
	}
}

func (g *GoodsServiceImpl) GoodsJingfenQuery(request *GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error) {
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsJingfenQueryResult
	err := g.service.Do(&res, GoodsJingfenQuery, param)
	return &res, err
}

func (g *GoodsServiceImpl) GoodsQuery(request *GoodsQueryRequest) (*GoodsQueryResult, error) {
	//goodsReqDTO
	param := map[string]interface{}{}
	param["goodsReqDTO"] = request
	var res GoodsQueryResult
	err := g.service.Do(&res, GoodsQuery, param)
	return &res, err
}

func (g *GoodsServiceImpl) GoodsPromotiongoodsinfoQuery(skuIds string) (*GoodsPromotiongoodsinfoResult, error) {
	//goodsReqDTO
	param := map[string]interface{}{}
	param["skuIds"] = skuIds
	var res GoodsPromotiongoodsinfoResult
	err := g.service.Do(&res, GoodsPromotiongoodsinfoQuery, param)
	return &res, err
}

func (g *GoodsServiceImpl) CategoryGoodsGetQuery(request *CategoryGoodsGetRequest) (*CategoryGoodsGetResult, error) {
	//goodsReqDTO
	param := map[string]interface{}{}
	param["req"] = request
	var res CategoryGoodsGetResult
	err := g.service.Do(&res, CategoryGoodsGetQuery, param)
	return &res, err
}

func (g *GoodsServiceImpl) GoodsGigfieldQuery(request *GoodsGigfieldRequest) (*GoodsGigfieldResult, error) {
	//goodsReqDTO
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsGigfieldResult
	err := g.service.Do(&res, GoodsGigfieldQuery, param)
	return &res, err
}
