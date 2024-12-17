package conf

import (
	"cupid-connector/internal/data"
	"cupid-connector/internal/model"
)

func UpdateBasic(bc model.BasicConf) error {
	data.Config.Basic = bc
	return saveEnv()
}

func UpdateMonitor(mc model.MonitorConf) error {
	data.Config.Monitor = mc
	return saveEnv()
}
