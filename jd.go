package jd

import (
	"encoding/json"
	"errors"
	"github.com/cliod/jd-go/common"
	"github.com/cliod/jd-go/log"
	"strings"
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
	// 可通过 Request 方法执行其他未封装的京东联盟API
	Request(v interface{}, param IParam) error

	// 参数校验
	CheckRequiredParameters(v interface{}) error

	// 获取解析正确结果
	GetResult(Result, error) ([]byte, error)
	// 解析并获取正确结果
	ParseResult(map[string]interface{}, ResponseKey) (map[string]interface{}, error)

	// 开启验证等级
	Validate(VCode)
	// 结果为map对象时的解封装等级，默认1
	// 0 不做解析，全部返回
	// 1 解析返回1层
	// 2 解析返回主要数据
	SetMapResultParseLevel(level int)
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

	otherService OtherService

	// 是否开启参数非空校验
	validateEmpty bool
	// 结果为map对象时的解封装等级，默认1
	mapResultParseLevel int
	// 请求的结果解封装等级 todo
	resultParseLevel int
}

// 默认京东联盟
func NewJdClient(appKey, secretKey, accessToken string) Service {
	return NewJdUnionService(appKey, secretKey, accessToken)
}

// 默认京东联盟
func NewJdService(appKey, secretKey string) Service {
	service := NewJdUnionService(appKey, secretKey, "")
	// 不做解析，兼容旧版本
	service.SetMapResultParseLevel(0)
	return service
}

func NewJdUnionService(appKey, secretKey, accessToken string) Service {
	return newService(appKey, secretKey, accessToken, UnionRootEndpoint)
}

func NewJosService(appKey, secretKey, accessToken string) Service {
	return newService(appKey, secretKey, accessToken, JosRootEndpoint)
}

func newService(appKey, secretKey, accessToken, routApi string) Service {
	impl := &ServiceImpl{}
	impl.SetConfig(newConfig(appKey, secretKey, accessToken, routApi))
	impl.SetHttpService(common.NewService())

	impl.goodsService = newGoodsService(impl)
	impl.promoteService = newPromotionService(impl)
	impl.couponService = newCouponService(impl)

	impl.giftService = newGiftService(impl)
	impl.orderService = newOrderService(impl)
	impl.activityService = newActivityService(impl)
	impl.positionService = newPositionService(impl)

	impl.otherService = newOtherService(impl.positionService, impl.activityService)

	impl.validateEmpty = false
	impl.mapResultParseLevel = 1
	return impl
}

func (ser *ServiceImpl) Get(url string, args interface{}) ([]byte, error) {
	return ser.service.Get(url, args)
}

func (ser *ServiceImpl) GetFor(v interface{}, url string, args *Param) error {
	res, err := ser.Get(url, args)
	if err != nil {
		log.Error("Request:", err)
		return err
	}
	return json.Unmarshal(res, v)
}

func (ser *ServiceImpl) GetGoodsService() GoodsService {
	return ser.goodsService
}

func (ser *ServiceImpl) GetCouponService() CouponService {
	return ser.couponService
}

func (ser *ServiceImpl) GetPromoteService() PromotionService {
	return ser.promoteService
}

func (ser *ServiceImpl) SetGoodsService(service GoodsService) {
	ser.goodsService = service
}

func (ser *ServiceImpl) SetCouponService(service CouponService) {
	ser.couponService = service
}

func (ser *ServiceImpl) SetPromoteService(service PromotionService) {
	ser.promoteService = service
}

func (ser *ServiceImpl) GetGiftService() GiftService {
	return ser.giftService
}

func (ser *ServiceImpl) GetOrderService() OrderService {
	return ser.orderService
}

func (ser *ServiceImpl) GetActivityService() ActivityService {
	return ser.activityService
}

func (ser *ServiceImpl) GetPositionService() PositionService {
	return ser.positionService
}

func (ser *ServiceImpl) SetGiftService(service GiftService) {
	ser.giftService = service
}

func (ser *ServiceImpl) SetActivityService(service ActivityService) {
	ser.activityService = service
}

func (ser *ServiceImpl) SetPositionService(service PositionService) {
	ser.positionService = service
}

func (ser *ServiceImpl) SetOrderService(service OrderService) {
	ser.orderService = service
}

// Deprecated: 使用新接口: GetPositionService 和 GetActivityService
func (ser *ServiceImpl) GetOtherService() OtherService {
	return ser.otherService
}

// Deprecated: 使用新接口: SetPositionService 和 SetActivityService
func (ser *ServiceImpl) SetOtherService(service OtherService) {
	ser.otherService = service
}

func (ser *ServiceImpl) GetConfig() *Config {
	return ser.config
}

func (ser *ServiceImpl) SetConfig(config *Config) {
	ser.config = config
}

func (ser *ServiceImpl) SetHttpService(service common.Service) {
	ser.service = service
}

func (ser *ServiceImpl) GetRouteApi() string {
	return ser.GetConfig().RouteApi
}

func (ser *ServiceImpl) SetRouteApi(api string) {
	ser.GetConfig().RouteApi = api
}

func (ser *ServiceImpl) Sign(method Method, param map[string]interface{}) (*Param, error) {
	c := ser.GetConfig()
	c.Method = method
	parameter := newParameter(c, param)
	parameter.attachSign()
	err := parameter.CheckRequiredParams()
	return NewParam(parameter), err
}

func (ser *ServiceImpl) Request(v interface{}, param IParam) error {
	var (
		method Method
		params map[string]interface{}
	)

	method = Method(param.Method())
	params = param.Params().ToMap()

	p, err := ser.Sign(method, params)
	if err != nil {
		log.Error("Sign:", err)
		return err
	}
	err = ser.GetFor(v, ser.GetConfig().RouteApi, p)
	if err != nil {
		log.Error("Result Err:", err)
	} else {
		log.Debug("Result:", v)
	}
	return err
}

// 非必填参数为空不会序列化
func (ser *ServiceImpl) CheckRequiredParameters(v interface{}) error {
	if !ser.validateEmpty {
		return nil
	}
	var param map[string]interface{}
	param, ok := v.(map[string]interface{})
	if !ok {
		bs, err := json.Marshal(v)
		if err != nil {
			log.Error("Validate Empty: ", err)
			return err
		}
		err = json.Unmarshal(bs, &param)
		if err != nil {
			log.Error("Validate Empty: ", err)
			return err
		}
	}
	for _, v := range param {
		if s, ok := v.(string); ok {
			if s == "" {
				err := errors.New("参数为空")
				log.Error("Validate Empty: ", err)
				return err
			}
		}
	}
	return nil
}

// Deprecated: 使用新接口: Request
func (ser *ServiceImpl) Do(v interface{}, method Method, params map[string]interface{}) error {
	param := NewTParam(method, params)
	return ser.Request(v, param)
}

func (ser *ServiceImpl) GetResult(res Result, err error) ([]byte, error) {
	if err != nil {
		log.Warn("Get Result: ", err)
		return nil, err
	}
	if res.IsSuccess() {
		return nil, res
	}
	return res.GetResult(), nil
}

func (ser *ServiceImpl) ParseResult(res map[string]interface{}, key ResponseKey) (map[string]interface{}, error) {
	if ser.mapResultParseLevel == 0 {
		return res, nil
	} else {
		//错误信息处理
		if res["error_response"] != nil {
			var errResponse *BaseResult
			errJson, err := json.Marshal(res)
			if err != nil {
				log.Error("Parse Result: ", err)
				return nil, err
			}
			if err := json.Unmarshal(errJson, &errResponse); err != nil {
				log.Error("Parse Result", err)
				return nil, err
			}
			return nil, errResponse
		}
		resp, ok := res[key.String()].(map[string]interface{})
		if ok {
			if ser.mapResultParseLevel == 1 {
				return resp, nil
			}
			var resName string
			if strings.Contains(key.String(), "response") {
				resName = "result"
			} else {
				resName = "queryResult"
			}
			str := resp[resName].(string)
			r := make(map[string]interface{})
			err := json.Unmarshal([]byte(str), &r)
			if err != nil {
				log.Error("Parse Result: ", err)
				return nil, err
			}
			return r, nil
		}
		return nil, errors.New("error parse result")
	}
}

func (ser *ServiceImpl) Validate(code VCode) {
	switch code {
	case NotEmpty:
		ser.validateEmpty = true
	case Non:
		ser.validateEmpty = false
	default:
	}
	//todo 其他校验
}

func (ser *ServiceImpl) SetMapResultParseLevel(level int) {
	ser.mapResultParseLevel = level
}
