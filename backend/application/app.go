package application

import (
	"context"
	"cupid-connector/backend/conf"
	"cupid-connector/backend/data"
	"cupid-connector/backend/ticker"
	"github.com/go-toast/toast"
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
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
	err := conf.InitEnv()
	if err != nil {
		log.Println(err)
		return
	}

	// 刷新信息
	a.RefreshInfo()

	// 自动登录
	autoLogin(a)

	// 初始化定时器
	err = ticker.Set()
	if err != nil {
		log.Println(err)
		return
	}

	// 启动监控
	go startMonitor(a)

	// 自动隐藏窗口
	if data.Config.Basic.AutoHide == "TRUE" {
		log.Println("自动隐藏窗口")
		n := toast.Notification{
			AppID:   AppID,
			Title:   "Cupid Connector",
			Message: "窗口已最小化到托盘",
		}
		err = n.Push()
		if err != nil {
			log.Println(err)
			return
		}
		runtime.WindowHide(ctx)
	}
}

func (a *App) shutdown(ctx context.Context) {
}
