package cmap

import "testing"

func TestCMap_ToStruct(t *testing.T) {
	type Te struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
	var te Te
	var param = New()
	param.Add("name", "tea")
	param.Add("id", "1")
	err := param.ToStruct(&te)
	if err != nil {
		t.Log(err)
		return
	}
	t.Log(te)

	t.Log(CMap{}.Add("id", "1").Set("test", "2"))
}
