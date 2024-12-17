package ticker

import (
	"cupid-connector/internal/conf"
	"log"
	"strconv"
	"time"
)

var Ticker *time.Ticker

// 初始化定时器
func Init() error {
	if Ticker != nil {
		log.Println("定时器已存在，停止定时器")
		Ticker.Stop()
	}
	if conf.Config.Monitor.Interval == "0" {
		log.Println("监控间隔为0，不启动定时器")
		return nil
	}
	interval, err := strconv.ParseFloat(conf.Config.Monitor.Interval, 64)
	if err != nil {
		log.Println("解析监控间隔失败！")
		return err
	}
	duration := time.Duration(int(interval)) * time.Minute
	Ticker = time.NewTicker(duration)
	log.Println("定时器已启动，间隔：" + duration.String())
	return nil
}
