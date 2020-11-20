package jd

import (
	"encoding/json"
	"fmt"
	"github.com/cliod/jd-go/common"
	"io/ioutil"
	"os"
	"testing"
)

// 获取商品列表
func TestServiceImpl0(t *testing.T) {
	c := GetConfig()
	j := NewJdService(c.AppKey, c.SecretKey)
	g := j.GetGoodsService()
	res, err := g.GoodsJingfenQuery(&GoodsJingfenQueryRequest{
		EliteId: 33,
	})
	t.Log("=======================商品查询=====================\n", res, err)
}

// 关键字获取商品（无权限）
func TestServiceImpl1(t *testing.T) {
	c := GetConfig()
	j := NewJdService(c.AppKey, c.SecretKey)
	g := j.GetGoodsService()
	res, err := g.GoodsQuery(&GoodsQueryRequest{})
	t.Log("=======================商品列表=====================\n", res, err)
}

// 商品类目 > fixme 无效签名??
func TestServiceImpl2(t *testing.T) {
	c := GetConfig()
	j := NewJdService(c.AppKey, c.SecretKey)
	g := j.GetGoodsService()

	res, err := g.CategoryGoodsGetQuery(&CategoryGoodsGetRequest{
		ParentId: 0,
		Grade:    0,
	})
	t.Log("=======================商品类目=====================\n", res, err)
}

// 获取商品列表
func TestServiceImpl3(t *testing.T) {
	c := GetConfig()
	j := NewJdService(c.AppKey, c.SecretKey)
	g := j.GetGoodsService()
	// 商品详情查询
	res, err := g.GoodsGigfieldQuery(&GoodsGigfieldRequest{
		SkuIds: []uint64{63682765730},
	})
	t.Log("=======================商品详情=====================\n", res, err)
}

// 获取商品详情
func TestServiceImpl4(t *testing.T) {
	c := GetConfig()
	j := NewJdService(c.AppKey, c.SecretKey)
	g := j.GetGoodsService()
	// 商品详情查询
	res, err := g.GoodsPromotiongoodsinfoQuery(
		"63682765730",
	)
	t.Log("=======================商品详情=====================\n", res, err)
}

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
		fmt.Println(err)
	}
	err = json.Unmarshal(b, &c)
	if err != nil {
		panic(err)
	}
	return c
}
