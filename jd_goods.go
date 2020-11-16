// jd联盟go-sdk
package jd

import "encoding/json"

type GoodsService interface {
	// 关键词商品查询接口【申请】
	GoodsQuery(request *GoodsQueryRequest) (*GoodsQueryResult, error)
}

type GoodsServiceImpl struct {
	service Service
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
	err = g.service.GetFor(&res, "")
	return &res, err
}
