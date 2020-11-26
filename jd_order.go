package jd

type OrderService interface {
	// 订单查询接口
	//    文档: https://union.jd.com/openplatform/api/10419
	OrderQuery(*OrderQueryRequest) (*OrderQueryResult, error)
	// 奖励订单查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11781
	OrderBonusQuery(*OrderBonusQueryRequest) (*OrderBonusQueryResult, error)
	// 订单行查询接口
	//    文档: https://union.jd.com/openplatform/api/12707
	OrderRowQuery(*OrderRowQueryRequest) (*OrderRowQueryResult, error)
}

type OrderServiceImpl struct {
	service Service
}

func newOrderService(service Service) OrderService {
	return &OrderServiceImpl{
		service: service,
	}
}

func (p *OrderServiceImpl) OrderQuery(request *OrderQueryRequest) (*OrderQueryResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderQueryResult
	err := p.service.Do(&res, OrderQuery, param)
	return &res, err
}

func (p *OrderServiceImpl) OrderBonusQuery(request *OrderBonusQueryRequest) (*OrderBonusQueryResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderBonusQueryResult
	err := p.service.Do(&res, OrderBonusQuery, param)
	return &res, err
}

func (p *OrderServiceImpl) OrderRowQuery(request *OrderRowQueryRequest) (*OrderRowQueryResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderRowQueryResult
	err := p.service.Do(&res, OrderRowQuery, param)
	return &res, err
}
