package jd

import (
	"encoding/json"
	"github.com/cliod/jd-go/common"
	"io/ioutil"
	"os"
	"testing"
)

func TestServiceImpl(t *testing.T) {

	c := GetConfig()

	j := NewJdService(c.AppKey, c.SecretKey)
	g := j.GetGoodsService()
	res, err := g.GoodsJingfenQuery(&GoodsJingfenQueryRequest{
		EliteId: 33,
	})
	t.Log(res, err)
	if err != nil {
		return
	}
	t.Log(res.GetResult(), err)
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
