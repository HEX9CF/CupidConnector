package main

import (
	"context"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/data"
	"cupid-connector/internal/ticker"
	"github.com/wailsapp/wails/v2/pkg/runtime"
	"log"
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

	// 自动登录
	a.autoLogin()

	// 初始化定时器
	err = ticker.Set()
	if err != nil {
		return
	}

	// 启动监控
	go a.startMonitor()

	// 自动隐藏窗口
	if data.Config.Basic.AutoHide == "TRUE" {
		log.Println("自动隐藏窗口")
		runtime.WindowHide(ctx)
	}
}

func (a *App) shutdown(ctx context.Context) {
}
