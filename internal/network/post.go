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
	"time"
)

const contentType = "application/x-www-form-urlencoded"

func PostRequest(url string, data string) (string, error) {
	log.Println("URL: ", url)
	log.Println("发送请求中...")
	// 创建一个带有超时设置的 HTTP 客户端
	client := &http.Client{
		Timeout: 5 * time.Second, // 设置超时时间为 5 秒
	}

	resp, err := client.Post(url, contentType, bytes.NewReader([]byte(data)))
	if err != nil {
		log.Println("请求发送失败！")
		return "", err
	}
	defer resp.Body.Close()

	log.Println("请求发送成功")

	if resp.StatusCode != http.StatusOK {
		log.Println("请求失败！")
		return "", errors.New("请求失败，" + resp.Proto + " " + resp.Status)
	}

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println("读取响应失败！")
		return "", err
	}
	bodyStr := string(body)

	return bodyStr, nil
}
