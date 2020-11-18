package jd

import (
	"encoding/json"
	"github.com/cliod/jd-go/common"
	"io/ioutil"
	"os"
	"testing"
	"time"
)

func TestServiceImpl(t *testing.T) {

	c := GetConfig()

	// jd api
	j := NewJdService(c.AppKey, c.SecretKey)

	// goods api
	g := j.GetGoodsService()
	res, err := g.GoodsJingfenQuery(&GoodsJingfenQueryRequest{
		EliteId: 33,
	})
	t.Log(res, err)
	res0, err := g.GoodsQuery(&GoodsQueryRequest{})
	t.Log(res0, err)
	res1, err := g.GoodsPromotiongoodsinfoQuery("63682765730")
	t.Log(res1, err)

	res2, err := g.CategoryGoodsGetQuery(&CategoryGoodsGetRequest{
		ParentId: 0,
		Grade:    0,
	})
	t.Log(res2, err)
	res3, err := g.GoodsGigfieldQuery(&GoodsGigfieldRequest{
		SkuIds: []uint64{63682765730},
	})
	t.Log(res3, err)

	// coupon api
	cc := j.GetCouponService()
	res4, err := cc.CouponQueryByList("https://coupon.m.jd.com/coupons/show.action?linkKey=AAROH_xIpeffAs_-naABEFoeRYiv3hJ69zyA-kQJMnHSr14qXhdrY5zeYZRSj6boZw2cJf-LDnD4b3HpA9b_Y_bjbeXiHA")
	t.Log(res4, err)

	o := j.GetOrderService()
	res5, err := o.OrderQuery(&OrderQueryRequest{
		PageNo:   1,
		PageSize: 30,
		TypeNum:  1,
		Time:     time.Now().Format("yyyyMMddHHmm"),
	})
	t.Log(res5, err)
}

func TestParameter_AttachSign(t *testing.T) {
	c := GetConfig()
	param := map[string]interface{}{}
	param["goodsReq"] = &GoodsJingfenQueryRequest{
		EliteId: 33,
	}
	p := NewParameter(NewConfig(c.AppKey, c.SecretKey), param)
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

type LocalConfig struct {
	AppKey    string `json:"app_key"`
	SecretKey string `json:"secret_key"`
}

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
		panic(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return c
}
