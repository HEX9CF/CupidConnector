package parser

import (
	"cupid-connector/backend/model"
	"encoding/json"
	"log"
	"strings"
)

func ParseLoginResp(bodyStr string) (model.LoginResp, error) {
	var resp model.LoginResp

	log.Println(bodyStr)
	bodyStr = strings.ReplaceAll(bodyStr, `'NOAUTH:*'`, `NOAUTH:*`)
	bodyStr = strings.ReplaceAll(bodyStr, "'", "\"")
	err := json.Unmarshal([]byte(bodyStr), &resp)
	if err != nil {
		log.Println("解析响应失败！")
		return model.LoginResp{}, err
	}

	return resp, nil
}
