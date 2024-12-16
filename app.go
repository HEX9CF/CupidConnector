package main

import (
	"context"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/model"
	"cupid-connector/internal/service"
	"github.com/go-toast/toast"
	"log"
	"os"
)

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
	err = service.Login()
	if err != nil {
		log.Println(err)
		notification := toast.Notification{
			AppID:   "Cupid Connector",
			Title:   "校园网登录失败",
			Message: "错误信息：" + err.Error(),
		}
		err := notification.Push()
		if err != nil {
			log.Fatalln(err)
			return
		}
		return
	}
	notification := toast.Notification{
		AppID:   "Cupid Connector",
		Title:   "校园网登录成功",
		Message: "登录成功，用户：" + conf.Config.Username + "，您已通过上网认证！",
	}
	err = notification.Push()
	if err != nil {
		log.Fatalln(err)
		return
	}
	if err == nil && conf.Config.AutoExit == "TRUE" {
		os.Exit(0)
	}
}

func (a *App) UpdateConf(config model.Conf) {
	err := service.UpdateConf(config)
	if err != nil {
		log.Println(err)
		n := toast.Notification{
			AppID:   "Cupid Connector",
			Title:   "配置更新失败",
			Message: "错误信息：" + err.Error(),
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	n := toast.Notification{
		AppID: "Cupid Connector",
		Title: "配置更新成功",
	}
	err = n.Push()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (a *App) GetConf() model.Resp {
	config := service.GetConf()
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
			AppID:   "Cupid Connector",
			Title:   "校园网登录失败",
			Message: "错误信息：" + err.Error(),
		}
		err := n.Push()
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	n := toast.Notification{
		AppID:   "Cupid Connector",
		Title:   "校园网登录成功",
		Message: "登录成功，用户：" + conf.Config.Username + "，您已通过上网认证！",
	}
	err = n.Push()
	if err != nil {
		log.Println(err)
		return
	}
	return
}

func (a *App) Logout() {
	err := service.Login()
	if err != nil {
		log.Println(err)
		n := toast.Notification{
			AppID:   "Cupid Connector",
			Title:   "校园网注销失败",
			Message: "错误信息：" + err.Error(),
		}
		err := n.Push()
		if err != nil {
			log.Println(err)
			return
		}
		return
	}
	n := toast.Notification{
		AppID: "Cupid Connector",
		Title: "校园网注销成功",
	}
	err = n.Push()
	if err != nil {
		log.Println(err)
		return
	}
	return
}
