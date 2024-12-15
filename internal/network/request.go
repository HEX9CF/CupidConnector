package network

import (
	"bytes"
	"io"
	"log"
	"net/http"
)

const contentType = "application/x-www-form-urlencoded"

func PostRequest(url string, data string) (string, error) {
	log.Println("URL: ", url)
	log.Println("发送请求中...")
	resp, err := http.Post(url, contentType, bytes.NewReader([]byte(data)))
	if err != nil {
		log.Println("请求发送失败！")
		return "", err
	}
	defer resp.Body.Close()

	log.Println("请求发送成功")

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败！")
		return "", err
	}
	bodyStr := string(body)

	return bodyStr, nil
}
