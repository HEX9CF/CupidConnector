package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
)

func Login() error {
	return api.Login(conf.Config.Url, conf.Config.Username, conf.Config.Password, "1")
}
