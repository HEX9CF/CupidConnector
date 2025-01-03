package ticker

import (
	"cupid-connector/backend/data"
	"log"
	"strconv"
	"time"
)

var Ticker *time.Ticker

// 初始化定时器
func Set() error {
	if data.Config.Monitor.Enable != "TRUE" || data.Config.Monitor.Interval == "0" {
		log.Println("关闭定时器")
		if Ticker != nil {
			Ticker.Stop()
		}
		return nil
	}
	interval, err := strconv.ParseFloat(data.Config.Monitor.Interval, 64)
	if err != nil {
		log.Println("解析监控间隔失败！")
		return err
	}
	duration := time.Duration(int(interval)) * time.Minute
	if Ticker == nil {
		Ticker = time.NewTicker(duration)
		log.Println("定时器已启动，间隔：" + duration.String())
	} else {
		Ticker.Reset(duration)
		log.Println("定时器已重置，间隔：" + duration.String())
	}
	return nil
}
