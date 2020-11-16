package common

import "encoding/json"

// 结构体转map[string]interface{}
func ToMap(request interface{}) map[string]interface{} {
	b, err := json.Marshal(request)
	if err != nil {
		return nil
	}
	var data map[string]interface{}
	err = json.Unmarshal(b, &data)
	if err != nil {
		return nil
	}
	return data
}
