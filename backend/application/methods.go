package application

import (
	"cupid-connector/backend/data"
	"cupid-connector/backend/model"
	"cupid-connector/backend/service/info"
	"cupid-connector/backend/service/login"
	"cupid-connector/backend/service/logout"
	"cupid-connector/backend/service/setting"
	"cupid-connector/backend/tray"
	"log"

	"github.com/go-toast/toast"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

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

func (a *App) GetInternetSpeed() model.Resp {
	return model.Resp{Code: model.ResponseCodeOk, Msg: "success", Data: data.InternetSpeed}
}

func (a *App) RefreshInternetSpeed() {
	runtime.EventsEmit(a.ctx, "updateInternetSpeed")
}

// 占位逼迫models生成Info类的，本身没用
func (a *App) GenerateInfo(model.Info) model.Resp {
	return model.Resp{Code: model.ResponseCodeError, Msg: "not implemented"}
}

// 占位逼迫models生成InternetSpeed类的，本身没用
func (a *App) GenerateInternetSpeed(speed model.InternetSpeed) model.Resp {
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
	runtime.Quit(a.ctx)
}
