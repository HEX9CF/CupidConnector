package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
	"errors"
)

func Login() error {
	if conf.Config.Basic.Username == "" || conf.Config.Basic.Password == "" {
		return errors.New("用户名或密码不能为空！")
	}
	if conf.Config.Basic.BaseUrl == "" {
		return errors.New("URL不能为空！")
	}
	return api.Login(conf.Config.Basic.BaseUrl+"/ac_portal/login.php",
		conf.Config.Basic.Username,
		conf.Config.Basic.Password,
		"1",
	)
}
