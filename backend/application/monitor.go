package application

import (
	"cupid-connector/backend/data"
	"cupid-connector/backend/ticker"
	"github.com/go-toast/toast"
	"log"
	"strconv"
)

var remainFlux float64

// 启动监控
func startMonitor(a *App) {
	for {
		if ticker.Ticker != nil {
			select {
			case <-ticker.Ticker.C:
				if data.Config.Monitor.Enable == "TRUE" &&
					data.Config.Monitor.Interval != "0" {
					log.Println("流量监控中...")

					// 自动重连
					autoReconnect(a)

					// 刷新信息
					refreshInfo(a)

					// 流量告警
					alert()

					// 自动注销
					autoExit(a)
				}
			}
		}
	}
}

// 刷新信息
func refreshInfo(a *App) {
	a.RefreshInfo()
	remainFlux = data.Info.Overall - data.Info.Used
}

// 自动重连
func autoReconnect(a *App) {
	if data.Config.Monitor.AutoReconnect == "TRUE" {
		log.Println("自动重连检测中...")
		if data.Info.Username == "" {
			a.Login()
		}
	}
}

// 流量告警
func alert() {
	if data.Info.Username == "" {
		return
	}
	if data.Config.Monitor.AlertThreshold != "0" {
		log.Println("流量告警检测中...")
		alertPercent, err := strconv.ParseFloat(data.Config.Monitor.AlertThreshold, 64)
		if err != nil {
			log.Println("解析警告阈值失败！")
			return
		}
		alertPercent *= 0.01
		if remainFlux/data.Info.Overall < alertPercent {
			msg := "当前已用流量 " + strconv.FormatFloat(data.Info.Used, 'f', 2, 64) + "MB，总流量 " + strconv.FormatFloat(data.Info.Overall, 'f', 2, 64) + "MB，剩余流量不足 " + strconv.FormatFloat(remainFlux, 'f', 2, 64) + "MB，已达到告警阈值（" + data.Config.Monitor.AlertThreshold + "%），请注意流量使用！"
			log.Println("流量告警：" + msg)
			n := toast.Notification{
				AppID:   AppID,
				Title:   "流量告警",
				Message: msg,
			}
			err = n.Push()
			if err != nil {
				log.Println(err)
				return
			}
		}
		log.Println("流量告警检测完成")
	}
}

// 自动注销
func autoExit(a *App) {
	if data.Info.Username == "" {
		return
	}
	if data.Config.Monitor.LogoutThreshold != "0" {
		log.Println("自动注销检测中...")
		logoutPercent, err := strconv.ParseFloat(data.Config.Monitor.LogoutThreshold, 64)
		if err != nil {
			log.Println("解析注销阈值失败！")
			return
		}
		logoutPercent *= 0.01
		if (data.Info.Overall-data.Info.Used)/data.Info.Overall < logoutPercent {
			msg := "已用流量 " + strconv.FormatFloat(data.Info.Used, 'f', 2, 64) + "MB，总流量 " + strconv.FormatFloat(data.Info.Overall, 'f', 2, 64) + "MB，剩余流量不足 " + strconv.FormatFloat(remainFlux, 'f', 2, 64) + "MB，已达到注销阈值（" + data.Config.Monitor.LogoutThreshold + "%），将自动注销校园网账号！"
			log.Println("自动注销：" + msg)
			n := toast.Notification{
				AppID:   AppID,
				Title:   "自动注销",
				Message: msg,
			}
			err = n.Push()
			if err != nil {
				log.Println(err)
			}
			a.Logout()
			refreshInfo(a)
		}
		log.Println("自动注销检测完成")
	}
}
