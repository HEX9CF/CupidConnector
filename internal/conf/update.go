package conf

import (
	"cupid-connector/internal/model"
)

func UpdateBasic(bc model.BasicConf) error {
	Config.Basic = bc
	return saveEnv()
}

func UpdateMonitor(mc model.MonitorConf) error {
	Config.Monitor = mc
	return saveEnv()
}
