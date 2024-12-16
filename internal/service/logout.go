package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
)

func Logout() error {
	return api.Logout(conf.Config.Url)
}
