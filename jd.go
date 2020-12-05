package jd

import (
	"encoding/json"
	"errors"
	"github.com/cliod/jd-go/common"
	"github.com/cliod/jd-go/log"
)

// 京东联盟API
//    文档: https://union.jd.com/openplatform/api
type Service interface {
	common.Service

	// 获取商品相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetGoodsService() GoodsService
	// 获取优惠卷相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetCouponService() CouponService
	// 获取推广相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetPromoteService() PromotionService
	// 获取礼金相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetGiftService() GiftService
	// 获取订单相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetOrderService() OrderService
	// 获取活动相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetActivityService() ActivityService
	// 获取推广位相关接口
	//    文档: https://union.jd.com/openplatform/api
	GetPositionService() PositionService
	// Deprecated: 弃用
	// 获取其他相关接口
	//    文档: https://union.jd.com/openplatform/api
	// 使用新接口: GetPositionService 和 GetActivityService
	GetOtherService() OtherService

	// 设置商品(自定义)相关接口
	SetGoodsService(GoodsService)
	// 设置优惠卷(自定义)相关接口
	SetCouponService(CouponService)
	// 设置推广(自定义)相关接口
	SetPromoteService(PromotionService)
	// 设置礼金(自定义)相关接口
	SetGiftService(GiftService)
	// 设置其他(自定义)相关接口
	SetActivityService(ActivityService)
	// 设置推广位(自定义)相关接口
	SetPositionService(PositionService)
	// 设置订单(自定义)相关接口
	SetOrderService(OrderService)
	// Deprecated: 弃用
	// 设置其他(自定义)相关接口
	// 使用新接口: SetPositionService 和 SetActivityService
	SetOtherService(OtherService)

	// 获取配置
	GetConfig() *Config
	// 设置配置
	SetConfig(config *Config)
	// 设置http请求
	SetHttpService(service common.Service)

	// 获取路由api
	GetRouteApi() string
	// 设置路由api
	SetRouteApi(string)

	// 执行http请求并解析结果
	// v 返回数据解析对象(指针)
	// method 请求路由方法
	// 业务参数
	// Deprecated: 使用新接口: Request
	Do(v interface{}, method Method, param map[string]interface{}) error

	// 执行http请求并解析结果
	// v 返回数据解析对象(指针)
	// method 请求路由方法
	// 业务参数
	Request(v interface{}, method Method, param map[string]interface{}) error
}

type ServiceImpl struct {
	service common.Service
	config  *Config

	goodsService    GoodsService
	promoteService  PromotionService
	couponService   CouponService
	giftService     GiftService
	orderService    OrderService
	activityService ActivityService
	positionService PositionService

	//Deprecated: 弃用
	otherService OtherService
}

// 默认京东联盟
func NewJdClient(appKey, secretKey string) Service {
	return NewJdUnionService(appKey, secretKey)
}

// 默认京东联盟
func NewJdService(appKey, secretKey string) Service {
	return NewJdUnionService(appKey, secretKey)
}

func NewJdUnionService(appKey, secretKey string) Service {
	return newService(appKey, secretKey, UnionRootEndpoint)
}

func NewJosService(appKey, secretKey string) Service {
	return newService(appKey, secretKey, JosRootEndpoint)
}

func newService(appKey, secretKey, routApi string) Service {
	impl := &ServiceImpl{}
	impl.SetConfig(newConfig(appKey, secretKey, routApi))
	impl.SetHttpService(common.NewService())

	impl.goodsService = newGoodsService(impl)
	impl.promoteService = newPromotionService(impl)
	impl.couponService = newCouponService(impl)

	impl.giftService = newGiftService(impl)
	impl.orderService = newOrderService(impl)
	impl.activityService = newActivityService(impl)
	impl.positionService = newPositionService(impl)

	impl.otherService = newOtherService(impl.positionService, impl.activityService)
	return impl
}

func (s *ServiceImpl) Get(url string, args interface{}) ([]byte, error) {
	return s.service.Get(url, args)
}

func (s *ServiceImpl) GetFor(v interface{}, url string, args *Param) error {
	res, err := s.Get(url, args)
	if err != nil {
		log.Error("Request:", err)
		return err
	}
	return json.Unmarshal(res, v)
}

func (s *ServiceImpl) GetGoodsService() GoodsService {
	return s.goodsService
}

func (s *ServiceImpl) GetCouponService() CouponService {
	return s.couponService
}

func (s *ServiceImpl) GetPromoteService() PromotionService {
	return s.promoteService
}

func (s *ServiceImpl) SetGoodsService(service GoodsService) {
	s.goodsService = service
}

func (s *ServiceImpl) SetCouponService(service CouponService) {
	s.couponService = service
}

func (s *ServiceImpl) SetPromoteService(service PromotionService) {
	s.promoteService = service
}

func (s *ServiceImpl) GetGiftService() GiftService {
	return s.giftService
}

func (s *ServiceImpl) GetOrderService() OrderService {
	return s.orderService
}

func (s *ServiceImpl) GetActivityService() ActivityService {
	return s.activityService
}

func (s *ServiceImpl) GetPositionService() PositionService {
	return s.positionService
}

func (s *ServiceImpl) SetGiftService(service GiftService) {
	s.giftService = service
}

func (s *ServiceImpl) SetActivityService(service ActivityService) {
	s.activityService = service
}

func (s *ServiceImpl) SetPositionService(service PositionService) {
	s.positionService = service
}

func (s *ServiceImpl) SetOrderService(service OrderService) {
	s.orderService = service
}

// Deprecated: 使用新接口: GetPositionService 和 GetActivityService
func (s *ServiceImpl) GetOtherService() OtherService {
	return s.otherService
}

// Deprecated: 使用新接口: SetPositionService 和 SetActivityService
func (s *ServiceImpl) SetOtherService(service OtherService) {
	s.otherService = service
}

func (s *ServiceImpl) GetConfig() *Config {
	return s.config
}

func (s *ServiceImpl) SetConfig(config *Config) {
	s.config = config
}

func (s *ServiceImpl) SetHttpService(service common.Service) {
	s.service = service
}

func (s *ServiceImpl) GetRouteApi() string {
	return s.GetConfig().RouteApi
}

func (s *ServiceImpl) SetRouteApi(api string) {
	s.GetConfig().RouteApi = api
}

func (s *ServiceImpl) Sign(method Method, param map[string]interface{}) (*Param, error) {
	c := s.GetConfig()
	c.Method = method
	parameter := newParameter(c, param)
	parameter.attachSign()
	err := parameter.CheckRequiredParams()
	return NewParam(parameter), err
}

func (s *ServiceImpl) Request(v interface{}, method Method, param map[string]interface{}) error {
	p, err := s.Sign(method, param)
	if err != nil {
		log.Error("Sign:", err)
		return err
	}
	err = s.GetFor(v, s.GetConfig().RouteApi, p)
	if err != nil {
		log.Error("Result Err:", err)
	} else {
		log.Debug("Result:", v)
	}
	return err
}

// 非必填参数为空不会序列化
func (s *ServiceImpl) CheckRequiredParameters(v interface{}) error {
	var param map[string]interface{}
	param, ok := v.(map[string]interface{})
	if !ok {
		bs, err := json.Marshal(v)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bs, &param)
		if err != nil {
			return err
		}
	}
	for _, v := range param {
		if s, ok := v.(string); ok {
			if s == "" {
				return errors.New("参数为空")
			}
		}
	}
	return nil
}

// Deprecated: 使用新接口: Request
func (s *ServiceImpl) Do(v interface{}, method Method, param map[string]interface{}) error {
	return s.Request(v, method, param)
}
