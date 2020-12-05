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

	// 订单查询接口
	//    文档: https://union.jd.com/openplatform/api/10419
	OrderQueryResult(*OrderQueryRequest) ([]byte, error)
	// 奖励订单查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11781
	OrderBonusQueryResult(*OrderBonusQueryRequest) ([]byte, error)
	// 订单行查询接口
	//    文档: https://union.jd.com/openplatform/api/12707
	OrderRowQueryResult(*OrderRowQueryRequest) ([]byte, error)

	// 订单查询接口
	//    文档: https://union.jd.com/openplatform/api/10419
	OrderQueryMap(*OrderQueryRequest) (map[string]interface{}, error)
	// 奖励订单查询接口【申请】
	//    文档: https://union.jd.com/openplatform/api/11781
	OrderBonusQueryMap(*OrderBonusQueryRequest) (map[string]interface{}, error)
	// 订单行查询接口
	//    文档: https://union.jd.com/openplatform/api/12707
	OrderRowQueryMap(*OrderRowQueryRequest) (map[string]interface{}, error)
}

type OrderServiceImpl struct {
	service Service
}

func newOrderService(service Service) OrderService {
	return &OrderServiceImpl{
		service: service,
	}
}

// 订单查询接口
//    文档: https://union.jd.com/openplatform/api/10419
func (order *OrderServiceImpl) OrderQuery(request *OrderQueryRequest) (*OrderQueryResult, error) {
	err := order.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderQueryResult
	err = order.service.Do(&res, OrderQuery, param)
	return &res, err
}

// 奖励订单查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11781
func (order *OrderServiceImpl) OrderBonusQuery(request *OrderBonusQueryRequest) (*OrderBonusQueryResult, error) {
	err := order.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderBonusQueryResult
	err = order.service.Do(&res, OrderBonusQuery, param)
	return &res, err
}

// 订单行查询接口
//    文档: https://union.jd.com/openplatform/api/12707
func (order *OrderServiceImpl) OrderRowQuery(request *OrderRowQueryRequest) (*OrderRowQueryResult, error) {
	err := order.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderRowQueryResult
	err = order.service.Do(&res, OrderRowQuery, param)
	return &res, err
}

// 订单查询接口
//    文档: https://union.jd.com/openplatform/api/10419
func (order *OrderServiceImpl) OrderQueryResult(request *OrderQueryRequest) ([]byte, error) {
	res, err := order.OrderQuery(request)
	return order.service.GetResult(res, err)
}

// 奖励订单查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11781
func (order *OrderServiceImpl) OrderBonusQueryResult(request *OrderBonusQueryRequest) ([]byte, error) {
	res, err := order.OrderBonusQuery(request)
	return order.service.GetResult(res, err)
}

// 订单行查询接口
//    文档: https://union.jd.com/openplatform/api/12707
func (order *OrderServiceImpl) OrderRowQueryResult(request *OrderRowQueryRequest) ([]byte, error) {
	res, err := order.OrderRowQuery(request)
	return order.service.GetResult(res, err)
}

// 订单查询接口
//    文档: https://union.jd.com/openplatform/api/10419
func (order *OrderServiceImpl) OrderQueryMap(request *OrderQueryRequest) (map[string]interface{}, error) {
	err := order.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res map[string]interface{}
	err = order.service.Do(&res, OrderQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if order.service.GetRouteApi() == JosRootEndpoint {
		key = OrderQueryResponce
	} else {
		key = OrderQueryResponse
	}
	return order.service.ParseResult(res, key)
}

// 奖励订单查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11781
func (order *OrderServiceImpl) OrderBonusQueryMap(request *OrderBonusQueryRequest) (map[string]interface{}, error) {
	err := order.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res map[string]interface{}
	err = order.service.Do(&res, OrderBonusQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if order.service.GetRouteApi() == JosRootEndpoint {
		key = OrderBonusQueryResponce
	} else {
		key = OrderBonusQueryResponse
	}
	return order.service.ParseResult(res, key)
}

// 订单行查询接口
//    文档: https://union.jd.com/openplatform/api/12707
func (order *OrderServiceImpl) OrderRowQueryMap(request *OrderRowQueryRequest) (map[string]interface{}, error) {
	err := order.service.CheckRequiredParameters(request)
	if err != nil {
		return nil, err
	}
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res map[string]interface{}
	err = order.service.Do(&res, OrderRowQuery, param)
	if err != nil {
		return nil, err
	}
	var key ResponseKey
	if order.service.GetRouteApi() == JosRootEndpoint {
		key = OrderRowQueryResponce
	} else {
		key = OrderRowQueryResponse
	}
	return order.service.ParseResult(res, key)
}
