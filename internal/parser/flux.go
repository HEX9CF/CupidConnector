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
			// End of the document, we're done
			return info
		case html.StartTagToken, html.SelfClosingTagToken:
			t := z.Token()
			if t.Data == "td" {
				// Read the text content of the td tag
				for {
					tt = z.Next()
					if tt == html.TextToken {
						text := string(z.Text())
						if key == "" {
							key = strings.TrimSpace(text)
						} else {
							value = strings.TrimSpace(text)
							// Map the key to the value in the info struct
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
