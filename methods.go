package main

import (
	"cupid-connector/internal/data"
	"cupid-connector/internal/model"
	"cupid-connector/internal/service/info"
	"cupid-connector/internal/service/login"
	"cupid-connector/internal/service/logout"
	"cupid-connector/internal/service/setting"
	"cupid-connector/internal/ticker"
	"cupid-connector/internal/tray"
	"log"
	"os"
	"strconv"

	"github.com/go-toast/toast"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// 更新基础配置
func (a *App) UpdateBasicConf(config model.BasicConf) {
	err := setting.UpdateBasicConf(config)
	if err != nil {
		log.Println(err)
		n := toast.Notification{
			AppID:   AppID,
			Title:   "配置更新失败",
			Message: "错误信息：" + err.Error(),
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		n := toast.Notification{
			AppID: AppID,
			Title: "配置更新成功",
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

// 更新监控配置
func (a *App) UpdateMonitorConf(config model.MonitorConf) {
	err := setting.UpdateMonitorConf(config)
	if err != nil {
		log.Println(err)
		n := toast.Notification{
			AppID:   AppID,
			Title:   "配置更新失败",
			Message: "错误信息：" + err.Error(),
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		n := toast.Notification{
			AppID: AppID,
			Title: "配置更新成功",
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	}
	err = ticker.Set()
	if err != nil {
		log.Println(err)
		return
	}
}

// 获取基础配置
func (a *App) GetBasicConf() model.Resp {
	config := setting.GetBasicConf()
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: config}
}

// 获取监控配置
func (a *App) GetMonitorConf() model.Resp {
	config := setting.GetMonitorConf()
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: config}
}

// 获取信息
func (a *App) GetInfo() model.Resp {
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: data.Info}
}

// 刷新流量信息
func (a *App) RefreshInfo() {
	log.Println("正在刷新流量信息...")
	err := info.Refresh()
	if err != nil {
		return
	}
	if tray.TrayApp.IsStartUp() {
		log.Println("更新托盘提示信息")
		tray.TrayApp.TooltipRefresh()
	}
	runtime.EventsEmit(a.ctx, "updateInfo")
}

// 占位逼迫models生成Info类的，本身没用
func (a *App) GenerateInfo(model.Info) model.Resp {
	return model.Resp{Code: model.ResponseCodeError, Msg: "not implemented"}
}

// 登录
func (a *App) Login() {
	err := login.Login()
	if err != nil {
		log.Println(err)
		n := toast.Notification{
			AppID:   AppID,
			Title:   "校园网登录失败",
			Message: "错误信息：" + err.Error(),
		}
		err := n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		n := toast.Notification{
			AppID:   AppID,
			Title:   "校园网登录成功",
			Message: "登录成功，用户：" + data.Config.Basic.Username + "，您已通过上网认证！",
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	}

	a.RefreshInfo()
}

// 注销
func (a *App) Logout() {
	err := logout.Logout()
	if err != nil {
		log.Println(err)
		n := toast.Notification{
			AppID:   AppID,
			Title:   "校园网注销失败",
			Message: "错误信息：" + err.Error(),
		}
		err := n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	} else {
		n := toast.Notification{
			AppID: AppID,
			Title: "校园网注销成功",
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	}

	a.RefreshInfo()
}

// 退出
func (a *App) Exit() {
	os.Exit(0)
}

// 启动监控
func (a *App) startMonitor() {
	for {
		if ticker.Ticker != nil {
			select {
			case <-ticker.Ticker.C:
				if data.Config.Monitor.Enable == "TRUE" &&
					data.Config.Monitor.Interval != "0" {
					log.Println("流量监控中...")

					// 刷新信息
					a.RefreshInfo()
					remainFlux := data.Info.Overall - data.Info.Used

					if data.Info.Username != "" {
						// 流量告警
						if data.Config.Monitor.AlertThreshold != "0" {
							log.Println("流量告警检测中...")
							alertPercent, err := strconv.ParseFloat(data.Config.Monitor.AlertThreshold, 64)
							if err != nil {
								log.Println("解析警告阈值失败！")
								continue
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
									continue
								}
							}
							log.Println("流量告警检测完成")
						}

						// 自动注销
						if data.Config.Monitor.LogoutThreshold != "0" {
							log.Println("自动注销检测中...")
							logoutPercent, err := strconv.ParseFloat(data.Config.Monitor.LogoutThreshold, 64)
							if err != nil {
								log.Println("解析注销阈值失败！")
								continue
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
									continue
								}
								err = logout.Logout()
								if err != nil {
									log.Println(err)
									continue
								}
								a.RefreshInfo()
							}
							log.Println("自动注销检测完成")
						}
					}
				}
			}
		}
	}
}
