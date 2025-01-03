package application

import (
	"cupid-connector/backend/data"
	"cupid-connector/backend/service/login"
	"github.com/go-toast/toast"
	"log"
	"time"
)

func autoLogin(a *App) {
	if data.Config.Basic.AutoLogin == "TRUE" {
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
		}

		// 刷新信息
		a.RefreshInfo()

		if data.Info.Username != "" && data.Config.Basic.AutoExit == "TRUE" {
			time.Sleep(3 * time.Second)
			a.Exit()
		}
	}
}
