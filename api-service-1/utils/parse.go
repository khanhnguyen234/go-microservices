package utils

import (
	"encoding/json"
)

func ParseStructToJson(value map[string]interface{}) string {
	jsonValue, _ := json.Marshal(value)
	return string(jsonValue)
}

func ParseJsonToStruct(value string) map[string]interface{} {
	var result map[string]interface{}
	json.Unmarshal([]byte(value), &result)
	return result
}
