package main

import (
	"context"
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/service"
	"cupid-connector/internal/ticker"
	"github.com/go-toast/toast"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
	"os"
	"strconv"
	"time"
)

const AppID = "Cupid Connector"

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
	err := conf.InitEnv()
	if err != nil {
		log.Println(err)
		return
	}
	if conf.Config.Basic.AutoLogin == "TRUE" {
		err = service.Login()
		if err != nil {
			log.Println(err)
			n := toast.Notification{
				AppID:   AppID,
				Title:   "校园网登录失败",
				Message: "错误信息：" + err.Error(),
			}
			err := n.Push()
			if err != nil {
				log.Fatalln(err)
			}
		} else {
			msg := "登录成功，用户：" + conf.Config.Basic.Username + "，您已通过上网认证！"
			if conf.Config.Basic.AutoExit == "TRUE" {
				msg += "程序将在3秒后自动退出"
			}
			n := toast.Notification{
				AppID:   AppID,
				Title:   "校园网登录成功",
				Message: msg,
			}
			err = n.Push()
			if err != nil {
				log.Fatalln(err)
			}
		}
		if err == nil && conf.Config.Basic.AutoExit == "TRUE" {
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	}

	// 初始化定时器
	err = ticker.Init()
	if err != nil {
		return
	}

	// 启动监控
	go startMonitor(ctx)
}

// 启动监控
func startMonitor(ctx context.Context) {
	for {
		if ticker.Ticker != nil {
			select {
			case <-ticker.Ticker.C:
				if conf.Config.Monitor.Enable == "TRUE" &&
					conf.Config.Monitor.Interval != "0" &&
					conf.Config.Basic.Username != "" {
					log.Println("流量监控中...")

					// 刷新信息
					runtime.EventsEmit(ctx, "refreshInfo")

					// 获取信息
					info, err := api.GetInfo(conf.Config.Basic.BaseUrl + "/ac_portal/userflux")
					if err != nil || info.Username == "" {
						log.Println("获取信息失败！")
						continue
					}
					remainValue := info.Overall - info.Used

					// 流量告警
					if conf.Config.Monitor.AlertThreshold != "0" {
						log.Println("流量告警检测中...")
						alertPercent, err := strconv.ParseFloat(conf.Config.Monitor.AlertThreshold, 64)
						if err != nil {
							log.Println("解析警告阈值失败！")
							continue
						}
						alertPercent *= 0.01
						if remainValue/info.Overall < alertPercent {
							msg := "当前已用流量 " + strconv.FormatFloat(info.Used, 'f', 2, 64) + "MB，总流量 " + strconv.FormatFloat(info.Overall, 'f', 2, 64) + "MB，剩余流量不足 " + strconv.FormatFloat(remainValue, 'f', 2, 64) + "MB，已达到告警阈值（" + conf.Config.Monitor.AlertThreshold + "%），请注意流量使用！"
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
					}

					// 自动注销
					if conf.Config.Monitor.LogoutThreshold != "0" {
						log.Println("自动注销检测中...")
						logoutPercent, err := strconv.ParseFloat(conf.Config.Monitor.LogoutThreshold, 64)
						if err != nil {
							log.Println("解析注销阈值失败！")
							continue
						}
						logoutPercent *= 0.01
						if (info.Overall-info.Used)/info.Overall < logoutPercent {
							err = service.Logout()
							if err != nil {
								log.Println(err)
								continue
							}
							msg := "已用流量 " + strconv.FormatFloat(info.Used, 'f', 2, 64) + "MB，总流量 " + strconv.FormatFloat(info.Overall, 'f', 2, 64) + "MB，剩余流量不足 " + strconv.FormatFloat(remainValue, 'f', 2, 64) + "MB，已达到注销阈值（" + conf.Config.Monitor.LogoutThreshold + "%），将自动注销校园网账号！"
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
						}
					}
				}
			}
		}
	}
}
