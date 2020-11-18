package jd

type OrderService interface {
	// 订单查询接口
	OrderQuery(*OrderQueryRequest) (*OrderQueryResult, error)
	// 奖励订单查询接口【申请】
	OrderBonusQuery(*OrderBonusRequest) (*OrderBonusResult, error)
	// 订单行查询接口
	OrderRowQuery(*OrderRowRequest) (*OrderRowResult, error)
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

func (p *OrderServiceImpl) OrderBonusQuery(request *OrderBonusRequest) (*OrderBonusResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderBonusResult
	err := p.service.Do(&res, OrderBonusQuery, param)
	return &res, err
}

func (p *OrderServiceImpl) OrderRowQuery(request *OrderRowRequest) (*OrderRowResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderRowResult
	err := p.service.Do(&res, OrderRowQuery, param)
	return &res, err
}
