package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
	"errors"
)

func Logout() error {
	if conf.Config.Basic.BaseUrl == "" {
		return errors.New("URL不能为空！")
	}
	return api.Logout(conf.Config.Basic.BaseUrl + "/ac_portal/login.php")
}
