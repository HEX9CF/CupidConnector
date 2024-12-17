package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
	"errors"
)

func Login() error {
	if conf.Config.Username == "" || conf.Config.Password == "" {
		return errors.New("用户名或密码不能为空！")
	}
	if conf.Config.BaseUrl == "" {
		return errors.New("URL不能为空！")
	}
	return api.Login(conf.Config.BaseUrl+"/ac_portal/login.php",
		conf.Config.Username,
		conf.Config.Password,
		"1",
	)
}
