package parser

import (
	"bufio"
	"bytes"
	"cupid-connector/internal/model"
	"golang.org/x/net/html"
	"log"
	"regexp"
	"strconv"
	"strings"
)

func ParseFluxInfo(body string) model.Info {
	if strings.Contains(body, "错误信息：临时用户和未登录用户不能查看流量信息！") {
		log.Println("错误信息：临时用户和未登录用户不能查看流量信息！")
		return model.Info{
			Username:       "临时用户",
			Overall:        1024,
			Used:           0,
			ExpirationTime: "未知",
			AccountStatus:  "未知",
		}
	}
	info := model.Info{}
	r := bytes.NewReader([]byte(body))
	z := html.NewTokenizer(bufio.NewReader(r))
	var (
		key   string
		value string
	)
	for {
		tt := z.Next()
		switch tt {
		case html.ErrorToken:
			return info
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "td" {
				for {
					tt = z.Next()
					if tt == html.TextToken {
						text := string(z.Text())
						if key == "" {
							key = strings.TrimSpace(text)
						} else {
							value = strings.TrimSpace(text)
							switch key {
							case "用户名称：":
								info.Username = value
							case "日流量额：":
								info.Overall = parseFluxValue(value)
							case "当天流量：":
								info.Used = parseFluxValue(value)
							case "过期时间：":
								info.ExpirationTime = value
							case "帐号状态：":
								info.AccountStatus = value
							}
							key = ""
							value = ""
						}
					} else if tt == html.EndTagToken && z.Token().Data == "td" {
						break
					}
				}
			}
		}
	}
	return info
}

func parseFluxValue(s string) float64 {
	var err error
	val := float64(0)

	reUnit := regexp.MustCompile(`\d*\.?\d*`)
	unit := reUnit.ReplaceAllString(s, "")

	reVal := regexp.MustCompile(`\d+(\.\d+)?`)
	valStr := reVal.FindString(s)
	if valStr == "" {
		val = 0
	} else {
		val, err = strconv.ParseFloat(valStr, 64)
		if err != nil {
			log.Println("解析流量值失败！")
			return 0
		}
	}

	if unit == "G" {
		val *= 1024
	}

	return val
}
