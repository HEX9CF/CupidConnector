package setting

import (
	"cupid-connector/internal/conf"
	"cupid-connector/internal/data"
	"cupid-connector/internal/model"
)

func UpdateBasicConf(c model.BasicConf) error {
	return conf.UpdateBasic(c)
}

func GetBasicConf() model.BasicConf {
	return data.Config.Basic
}

func UpdateMonitorConf(c model.MonitorConf) error {
	return conf.UpdateMonitor(c)
}

func GetMonitorConf() model.MonitorConf {
	return data.Config.Monitor
}
