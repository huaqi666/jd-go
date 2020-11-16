package jd

import (
	"testing"
)

func TestServiceImpl(t *testing.T) {
	t.Logf("1")

	j := NewJdService("<app_key>", "<app_secret>")
	g := j.GetGoodsService()
	t.Log(g.GoodsJingfenQuery(nil))
}
