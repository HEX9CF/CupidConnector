package network

import (
	"cupid-connector/internal/model"
	"encoding/json"
	"log"
	"strings"
)

func ResolveResponse(bodyStr string) (model.Response, error) {
	var resp model.Response

	//log.Println(bodyStr)
	bodyStr = strings.ReplaceAll(bodyStr, "'", "\"")
	err := json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		log.Println("解析响应失败！")
		return model.Response{}, err
	}

	return resp, nil
}
