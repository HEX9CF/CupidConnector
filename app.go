package main

import (
	"context"
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/model"
	"cupid-connector/internal/service"
	"github.com/go-toast/toast"
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
	go a.startMonitor()
}

func (a *App) UpdateBasicConf(config model.BasicConf) {
	err := service.UpdateBasicConf(config)
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

func (a *App) UpdateMonitorConf(config model.MonitorConf) {
	err := service.UpdateMonitorConf(config)
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

func (a *App) GetBasicConf() model.Resp {
	config := service.GetBasicConf()
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: config}
}

func (a *App) GetMonitorConf() model.Resp {
	config := service.GetMonitorConf()
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: config}
}

func (a *App) GetInfo() model.Resp {
	info, err := service.GetInfo()
	if err != nil {
		return model.Resp{Code: model.ResponseCodeError, Msg: err.Error()}
	}
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: info}
}

// UpdateInfo拿来占位逼迫models生成Info类的，本身没用
func (a *App) UpdateInfo(model.Info) model.Resp {
	return model.Resp{Code: model.ResponseCodeError, Msg: "not implemented"}
}

func (a *App) Login() {
	err := service.Login()
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
			Message: "登录成功，用户：" + conf.Config.Basic.Username + "，您已通过上网认证！",
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func (a *App) Logout() {
	err := service.Logout()
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
}

func (a *App) Exit() {
	os.Exit(0)
}

func (a *App) startMonitor() {
	interval, err := strconv.ParseFloat(conf.Config.Monitor.Interval, 64)
	if err != nil {
		log.Println("解析监控间隔失败！")
		interval = 5
	}
	ticker := time.NewTicker(time.Duration(int(interval)) * time.Minute)
	defer ticker.Stop()

	for {
		select {
		case <-ticker.C:
			if conf.Config.Monitor.Enable == "TRUE" {
				info, err := api.GetInfo(conf.Config.Basic.BaseUrl + "/ac_portal/userflux")
				if err != nil || info.Username == "" {
					log.Println("获取信息失败！")
					continue
				}

				if conf.Config.Monitor.AlertThreshold != "0" {
					alertPercent, err := strconv.ParseFloat(conf.Config.Monitor.AlertThreshold, 64)
					if err != nil {
						log.Println("解析警告阈值失败！")
						continue
					}
					alertPercent *= 0.01
				}
			}
		}
	}
}
