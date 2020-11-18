package jd

import (
	"encoding/json"
	"github.com/cliod/jd-go/common"
)

type Service interface {
	common.Service

	// 获取商品相关接口
	GetGoodsService() GoodsService
	// 获取优惠卷相关接口
	GetCouponService() CouponService
	// 获取推广相关接口
	GetPromoteService() PromotionService
	// 获取礼金相关接口
	GetGiftService() GiftService
	// 获取订单相关接口
	GetOrderService() OrderService
	// 获取其他相关接口
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
	SetOtherService(OtherService)
	// 设置订单(自定义)相关接口
	SetOrderService(OrderService)

	// 获取配置
	GetConfig() *Config
	// 设置配置
	SetConfig(config *Config)
	// 设置http请求
	SetHttpService(service common.Service)

	// 执行Get操作
	Do(v interface{}, method Method, param map[string]interface{}) error
}

type ServiceImpl struct {
	service common.Service
	config  *Config

	goodsService   GoodsService
	promoteService PromotionService
	couponService  CouponService
	giftService    GiftService
	orderService   OrderService
	otherService   OtherService
}

func NewJdService(appKet, secretKey string) Service {
	impl := &ServiceImpl{}
	impl.SetConfig(NewConfig(appKet, secretKey))
	impl.SetHttpService(common.NewService())

	impl.goodsService = newGoodsService(impl)
	impl.promoteService = newPromotionService(impl)
	impl.couponService = newCouponService(impl)

	impl.giftService = newGiftService(impl)
	impl.orderService = newOrderService(impl)
	impl.otherService = newOtherService(impl)
	return impl
}

func (s *ServiceImpl) Get(url string, args interface{}) ([]byte, error) {
	return s.service.Get(url, args)
}

func (s *ServiceImpl) GetFor(v interface{}, url string, args *Param) error {
	res, err := s.Get(url, args)
	if err != nil {
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

func (s *ServiceImpl) GetOtherService() OtherService {
	return s.otherService
}

func (s *ServiceImpl) SetGiftService(service GiftService) {
	s.giftService = service
}

func (s *ServiceImpl) SetOtherService(service OtherService) {
	s.otherService = service
}

func (s *ServiceImpl) SetOrderService(service OrderService) {
	s.orderService = service
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

func (s *ServiceImpl) Sign(method Method, param map[string]interface{}) *Param {
	c := s.GetConfig()
	c.Method = method
	parameter := NewParameter(c, param)
	parameter.attachSign()
	return &Param{
		ParamJson: parameter.getParamString(),
		Config: Config{
			AppKey:      parameter.AppKey,
			Method:      parameter.Method,
			AccessToken: parameter.AccessToken,
			Timestamp:   parameter.Timestamp,
			Format:      parameter.Format,
			Version:     parameter.Version,
			SignMethod:  parameter.SignMethod,
			Sign:        parameter.Sign,
		},
	}
}

func (s *ServiceImpl) Do(v interface{}, method Method, param map[string]interface{}) error {
	return s.GetFor(v, BaseUrl, s.Sign(method, param))
}
