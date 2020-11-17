// jd联盟go-sdk
package jd

type GoodsService interface {
	// 京粉精选商品查询接口
	GoodsJingfenQuery(*GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error)
	// 关键词商品查询接口【申请】
	GoodsQuery(*GoodsQueryRequest) (*GoodsQueryResult, error)
}

type GoodsServiceImpl struct {
	service Service
}

func newGoodsService(service Service) GoodsService {
	return &GoodsServiceImpl{
		service: service,
	}
}

func (g *GoodsServiceImpl) GoodsQuery(request *GoodsQueryRequest) (*GoodsQueryResult, error) {
	//goodsReqDTO
	param := map[string]interface{}{}
	param["goodsReqDTO"] = request
	var res GoodsQueryResult
	err := g.service.Do(&res, GoodsQuery, param)
	return &res, err
}

func (g *GoodsServiceImpl) GoodsJingfenQuery(request *GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error) {
	param := map[string]interface{}{}
	param["goodsReq"] = request
	var res GoodsJingfenQueryResult
	err := g.service.Do(&res, GoodsJingfenQuery, param)
	return &res, err
}
