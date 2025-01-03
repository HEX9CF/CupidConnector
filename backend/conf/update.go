package conf

import (
	"cupid-connector/backend/data"
	"cupid-connector/backend/model"
)

func UpdateBasic(bc model.BasicConf) error {
	data.Config.Basic = bc
	return saveEnv()
}

func UpdateMonitor(mc model.MonitorConf) error {
	data.Config.Monitor = mc
	return saveEnv()
}
