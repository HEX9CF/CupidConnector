package main

import (
	"context"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/data"
	"cupid-connector/internal/service/login"
	"cupid-connector/internal/ticker"
	"github.com/go-toast/toast"
	"log"
	"os"
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
	// 刷新信息
	a.RefreshInfo()
	if data.Config.Basic.AutoLogin == "TRUE" {
		err = login.Login()
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
			msg := "登录成功，用户：" + data.Config.Basic.Username + "，您已通过上网认证！"
			if data.Config.Basic.AutoExit == "TRUE" {
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
			// 刷新信息
			a.RefreshInfo()
		}
		if err == nil && data.Config.Basic.AutoExit == "TRUE" {
			time.Sleep(3 * time.Second)
			os.Exit(0)
		}
	}

	// 初始化定时器
	err = ticker.Set()
	if err != nil {
		return
	}

	// 启动监控
	go a.startMonitor()
}
