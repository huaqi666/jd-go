package jd

import (
	"encoding/json"
	"fmt"
	"github.com/cliod/jd-go/common"
	"io/ioutil"
	"os"
	"testing"
)

type LocalConfig struct {
	AppKey    string `json:"app_key"`
	SecretKey string `json:"secret_key"`
}

// 获取测试配置信息
func GetConfig() LocalConfig {
	var c LocalConfig
	f, err := os.Open("./config.json")
	if err != nil {
		f, err := os.Create("./config.json")
		if err != nil {
			panic(err)
		}
		b, err := json.Marshal(c)
		if err != nil {
			panic(err)
		}
		_, _ = f.Write(b)
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return c
}

// 签名测试
func TestParameter_AttachSign(t *testing.T) {
	c := GetConfig()
	param := map[string]interface{}{}
	param["goodsReq"] = &GoodsJingfenQueryRequest{
		EliteId: 33,
	}
	p := newParameter(NewConfig(c.AppKey, c.SecretKey), param)
	p.Method = "jd.union.open.goods.jingfen.query"
	p.attachSign()

	r, _ := json.Marshal(param)
	b, err := common.NewService().Get(BaseUrl, Param{
		ParamJson: string(r),
		Config: Config{
			AppKey:      p.AppKey,
			Method:      p.Method,
			AccessToken: p.AccessToken,
			Timestamp:   p.Timestamp,
			Format:      p.Format,
			Version:     p.Version,
			SignMethod:  p.SignMethod,
			Sign:        p.Sign,
		},
	})
	t.Log(string(b), err)
}

// 优惠券领取情况查询接口【申请】
func TestCouponServiceImpl_CouponQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	couponService := service.GetCouponService()

	res, err := couponService.CouponQueryByList("https://coupon.m.jd.com/coupons/show.action?linkKey=AAROH_xIpeffAs_-naABEFoe_GDEC4dVfVg03Fj3h8FXrH2uMMS7Y_UectztQMsMX3S6kCZX2ZFNkNK1ppUHnuinop15Fw")
	t.Log(res, err)
}

// 礼金创建
func TestGiftServiceImpl_CouponGiftGet(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	giftService := service.GetGiftService()

	res, err := giftService.CouponGiftGet(&CouponGiftGetRequest{
		SkuMaterialId:    "",
		Discount:         0,
		Amount:           0,
		ReceiveStartTime: "",
		ReceiveEndTime:   "",
		IsSpu:            0,
		ExpireType:       0,
		Share:            0,
	})
	t.Log(res, err)
}

// 礼金停止
func TestGiftServiceImpl_CouponGiftStop(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	giftService := service.GetGiftService()

	res, err := giftService.CouponGiftStopBy("")
	t.Log(res, err)
}

// 礼金效果数据
func TestGiftServiceImpl_GiftStatisticCouponQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	giftService := service.GetGiftService()

	res, err := giftService.StatisticGiftCouponQuery(&StatisticGiftCouponQueryRequest{
		SkuId:         0,
		GiftCouponKey: "",
		StartTime:     "",
		CreateTime:    "",
	})
	t.Log(res, err)
}

func TestGoodsServiceImpl_CategoryGoodsGetQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	goodsService := service.GetGoodsService()

	res, err := goodsService.CategoryGoodsGetQuery(&CategoryGoodsGetRequest{
		ParentId: 0,
		Grade:    0,
	})
	t.Log("=======================商品类目=====================\n", res, err)
}

// 获取商品列表
func TestGoodsServiceImpl_GoodsGigfieldQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	goodsService := service.GetGoodsService()
	// 商品详情查询
	res, err := goodsService.GoodsGigfieldQuery(&GoodsGigFieldQueryRequest{
		SkuIds: []uint64{63682765730},
	})
	t.Log(res, err)
}

// 获取商品列表
func TestGoodsServiceImpl_GoodsJingfenQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	goodsService := service.GetGoodsService()
	res, err := goodsService.GoodsJingfenQuery(&GoodsJingfenQueryRequest{
		EliteId: 33,
	})
	t.Log(res, err)
}

// 获取商品详情
func TestGoodsServiceImpl_GoodsPromotiongoodsinfoQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	goodsService := service.GetGoodsService()
	// 商品详情查询
	res, err := goodsService.GoodsPromotiongoodsinfoQuery(
		"63682765730",
	)
	t.Log(res, err)
}

// 关键字获取商品
func TestGoodsServiceImpl_GoodsQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	goodsService := service.GetGoodsService()
	res, err := goodsService.GoodsQuery(&GoodsQueryRequest{})
	t.Log(res, err)
}

func TestOrderServiceImpl_OrderBonusQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	orderService := service.GetOrderService()

	res, err := orderService.OrderBonusQuery(&OrderBonusQueryRequest{})
	t.Log(res, err)
}

func TestOrderServiceImpl_OrderQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	orderService := service.GetOrderService()

	res, err := orderService.OrderQuery(&OrderQueryRequest{})
	t.Log(res, err)
}

func TestOrderServiceImpl_OrderRowQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	orderService := service.GetOrderService()

	res, err := orderService.OrderRowQuery(&OrderRowQueryRequest{})
	t.Log(res, err)
}

func TestOtherServiceImpl_ActivityQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	otherService := service.GetOtherService()

	res, err := otherService.ActivityQuery(&ActivityQueryRequest{})
	t.Log(res, err)
}

func TestOtherServiceImpl_PositionCreate(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	otherService := service.GetOtherService()

	res, err := otherService.PositionCreate(&PositionCreateRequest{})
	t.Log(res, err)
}

func TestOtherServiceImpl_PositionQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	otherService := service.GetOtherService()

	res, err := otherService.PositionQuery(&PositionQueryRequest{})
	t.Log(res, err)
}

func TestOtherServiceImpl_StatisticsRedpacketQuery(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	otherService := service.GetOtherService()

	res, err := otherService.StatisticsRedpacketQuery(&StatisticsRedpacketQueryRequest{})
	t.Log(res, err)
}

func TestOtherServiceImpl_UserPidGet(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	otherService := service.GetOtherService()

	res, err := otherService.UserPidGet(&UserPidGetRequest{})
	t.Log(res, err)
}

func TestPromotionServiceImpl_PromotionBysubunionidGet(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	promoteService := service.GetPromoteService()

	res, err := promoteService.PromotionBysubunionidGet(&PromotionBysubunionidGetRequest{})
	t.Log(res, err)
}

func TestPromotionServiceImpl_PromotionByunionidGet(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	promoteService := service.GetPromoteService()

	res, err := promoteService.PromotionByunionidGet(&PromotionByunionidGetRequest{})
	t.Log(res, err)
}

func TestPromotionServiceImpl_PromotionCommonGet(t *testing.T) {
	config := GetConfig()
	service := NewJdService(config.AppKey, config.SecretKey)
	promoteService := service.GetPromoteService()

	res, err := promoteService.PromotionCommonGet(&PromoteCommonGetRequest{})
	t.Log(res, err)
}
