package conf

import (
	"cupid-connector/internal/model"
	"os"
)

var (
	Config model.Conf
)

func getEnv() {
	Config.BaseUrl = os.Getenv("BASE_URL")
	Config.Username = os.Getenv("STU_USERNAME")
	Config.Password = os.Getenv("STU_PASSWORD")
	Config.AutoLogin = os.Getenv("AUTO_LOGIN")
	Config.AutoExit = os.Getenv("AUTO_EXIT")
}
