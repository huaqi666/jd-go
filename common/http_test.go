package common

import (
	"testing"
)

func TestServiceImpl_Get(t *testing.T) {
	t.Log(NewService().Get("<url>", nil))
}
