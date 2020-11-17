package common

import "testing"

func TestMd5ToUpper(t *testing.T) {
	t.Log(Md5ToUpper("app_key%sformatjsonmethodjd.union.open.goods.jingfen." +
		"queryparam_json{\"goodsReq\":{\"eliteId\":\"33\"}}sign_methodmd5timestamp2020-11-17 13:20:47v1.0%s"))
}
