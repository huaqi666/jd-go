// jd联盟go-sdk
package jd

import "encoding/json"

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
	param := map[string]string{}
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	param["goodsReqDTO"] = string(b)
	var res GoodsQueryResult
	err = g.service.DoGet(&res, GoodsQuery, param)
	return &res, err
}

func (g *GoodsServiceImpl) GoodsJingfenQuery(request *GoodsJingfenQueryRequest) (*GoodsJingfenQueryResult, error) {
	param := map[string]string{}
	b, err := json.Marshal(request)
	if err != nil {
		return nil, err
	}
	param["goodsReq"] = string(b)
	var res GoodsJingfenQueryResult
	err = g.service.DoGet(&res, GoodsJingfenQuery, param)
	return &res, err
}
