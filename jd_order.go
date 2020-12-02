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
func (p *OrderServiceImpl) OrderQuery(request *OrderQueryRequest) (*OrderQueryResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderQueryResult
	err := p.service.Do(&res, OrderQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 奖励订单查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11781
func (p *OrderServiceImpl) OrderBonusQuery(request *OrderBonusQueryRequest) (*OrderBonusQueryResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderBonusQueryResult
	err := p.service.Do(&res, OrderBonusQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 订单行查询接口
//    文档: https://union.jd.com/openplatform/api/12707
func (p *OrderServiceImpl) OrderRowQuery(request *OrderRowQueryRequest) (*OrderRowQueryResult, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res OrderRowQueryResult
	err := p.service.Do(&res, OrderRowQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return &res, err
}

// 订单查询接口
//    文档: https://union.jd.com/openplatform/api/10419
func (p *OrderServiceImpl) OrderQueryResult(request *OrderQueryRequest) ([]byte, error) {
	res, err := p.OrderQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 奖励订单查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11781
func (p *OrderServiceImpl) OrderBonusQueryResult(request *OrderBonusQueryRequest) ([]byte, error) {
	res, err := p.OrderBonusQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 订单行查询接口
//    文档: https://union.jd.com/openplatform/api/12707
func (p *OrderServiceImpl) OrderRowQueryResult(request *OrderRowQueryRequest) ([]byte, error) {
	res, err := p.OrderRowQuery(request)
	if err != nil {
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

// 订单查询接口
//    文档: https://union.jd.com/openplatform/api/10419
func (p *OrderServiceImpl) OrderQueryMap(request *OrderQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, OrderQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}

// 奖励订单查询接口【申请】
//    文档: https://union.jd.com/openplatform/api/11781
func (p *OrderServiceImpl) OrderBonusQueryMap(request *OrderBonusQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, OrderBonusQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}

// 订单行查询接口
//    文档: https://union.jd.com/openplatform/api/12707
func (p *OrderServiceImpl) OrderRowQueryMap(request *OrderRowQueryRequest) (map[string]interface{}, error) {
	param := map[string]interface{}{}
	param["orderReq"] = request
	var res map[string]interface{}
	err := p.service.Do(&res, OrderRowQuery, param)
	debugLog.Println(LogPrefix, "Result: ", res)
	return res, err
}
