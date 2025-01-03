package application

import (
	"cupid-connector/backend/model"
	"cupid-connector/backend/service/setting"
	"cupid-connector/backend/ticker"
	"github.com/go-toast/toast"
	"log"
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
