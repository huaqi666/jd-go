package jd_method

import (
	"testing"
)

func TestMethodOf(t *testing.T) {
	t.Log(MethodOf(KeyOfResponce(GoodsGigfieldQuery)))
}

func TestKeyOfResponse(t *testing.T) {
	t.Log(KeyOfResponse(GoodsGigfieldQuery))
}

func TestKeyOfResponce(t *testing.T) {
	t.Log(KeyOfResponce(GoodsGigfieldQuery))
}
