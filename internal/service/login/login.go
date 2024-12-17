package login

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/data"
	"errors"
)

func Login() error {
	if data.Config.Basic.Username == "" || data.Config.Basic.Password == "" {
		return errors.New("用户名或密码不能为空！")
	}
	if data.Config.Basic.BaseUrl == "" {
		return errors.New("URL不能为空！")
	}
	return api.Login(data.Config.Basic.BaseUrl+"/ac_portal/login.php",
		data.Config.Basic.Username,
		data.Config.Basic.Password,
		"1",
	)
}
