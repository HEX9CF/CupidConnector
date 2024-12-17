package logout

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/data"
	"errors"
)

func Logout() error {
	if data.Config.Basic.BaseUrl == "" {
		return errors.New("URL不能为空！")
	}
	return api.Logout(data.Config.Basic.BaseUrl + "/ac_portal/login.php")
}
