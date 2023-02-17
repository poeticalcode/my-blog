package util

import (
	"encoding/json"
	"io"
	"net/http"
)

// ParseResponse 解析响应结果
func ParseResponse(response *http.Response) (map[string]interface{}, error) {
	var result map[string]interface{}
	body, err := io.ReadAll(response.Body)
	if err == nil {
		// 将 JSON 数据转为 map
		err = json.Unmarshal(body, &result)
	}
	return result, err
}
