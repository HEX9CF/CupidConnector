package conf

import (
	"CupidConnector/internal/model"
	"os"
)

var (
	Config model.Conf
)

func getEnv() {
	Config.Url = os.Getenv("STU_URL")
	Config.Username = os.Getenv("STU_USERNAME")
	Config.Password = os.Getenv("STU_PASSWORD")
	Config.AutoExit = os.Getenv("AUTO_EXIT")
}
