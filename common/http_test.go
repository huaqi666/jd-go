package common

import (
	"testing"
)

func TestServiceImpl_Get(t *testing.T) {
	res, err := NewService().Get("https://www.baidu.com", nil)
	t.Log(string(res), err)
}
