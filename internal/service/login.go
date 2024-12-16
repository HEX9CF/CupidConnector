package service

import (
	"CupidConnector/internal/api"
	"CupidConnector/internal/conf"
)

func Login() error {
	return api.Login(conf.Config.Url, conf.Config.Username, conf.Config.Password, "1")
}
