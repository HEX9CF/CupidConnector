/*
// @author Howard Zheng <hex9cf@aliyun.com>
*/
package network

import (
	"bytes"
	"errors"
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

	if resp.StatusCode != http.StatusOK {
		log.Println("请求失败！")
		return "", errors.New("请求失败，状态码：" + string(rune(resp.StatusCode)))
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败！")
		return "", err
	}
	bodyStr := string(body)

	return bodyStr, nil
}
