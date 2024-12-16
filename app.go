package main

import (
	"context"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/model"
	"cupid-connector/internal/service"
	"log"
	"os"

	"github.com/go-toast/toast"
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
		return
	}
	err = service.Login()
	if err != nil {
		notification := toast.Notification{
			AppID:   "Cupid Connector",
			Title:   "校园网登录失败",
			Message: "错误信息：" + err.Error(),
		}
		err := notification.Push()
		if err != nil {
			log.Fatalln(err)
		}
	} else {
		notification := toast.Notification{
			AppID:   "Cupid Connector",
			Title:   "校园网登录成功",
			Message: "登录成功，用户：" + conf.Config.Username + "，您已通过上网认证！",
		}
		err := notification.Push()
		if err != nil {
			log.Fatalln(err)
		}
	}
	if err == nil && conf.Config.AutoExit == "TRUE" {
		os.Exit(0)
	}
}

func (a *App) UpdateConf(config model.Conf) model.Resp {
	err := service.UpdateConf(config)
	if err != nil {
		return model.Resp{Code: model.ResponseCodeError, Msg: err.Error()}
	}
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success"}
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

func (a *App) UpdateInfo(model.Info) model.Resp {
	return model.Resp{Code: model.ResponseCodeError, Msg: "not implemented"}
}

func (a *App) Login() model.Resp {
	err := service.Login()
	if err != nil {
		n := toast.Notification{
			AppID:   "Cupid Connector",
			Title:   "校园网登录失败",
			Message: "错误信息：" + err.Error(),
		}
		err := n.Push()
		if err != nil {
			return model.Resp{Code: model.ResponseCodeError, Msg: err.Error()}
		}
	}
	n := toast.Notification{
		AppID:   "Cupid Connector",
		Title:   "校园网登录成功",
		Message: "登录成功，用户：" + conf.Config.Username + "，您已通过上网认证！",
	}
	err = n.Push()
	if err != nil {
		return model.Resp{Code: model.ResponseCodeError, Msg: err.Error()}
	}
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success"}
}
