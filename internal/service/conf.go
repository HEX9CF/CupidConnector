package service

import (
	"cupid-connector/internal/conf"
	"cupid-connector/internal/model"
)

func UpdateBasicConf(c model.BasicConf) error {
	return conf.UpdateBasic(c)
}

func GetBasicConf() model.BasicConf {
	return conf.Config.Basic
}

func UpdateMonitorConf(c model.MonitorConf) error {
	return conf.UpdateMonitor(c)
}

func GetMonitorConf() model.MonitorConf {
	return conf.Config.Monitor
}
