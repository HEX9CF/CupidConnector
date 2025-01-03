package logout

import (
	"cupid-connector/backend/api"
	"cupid-connector/backend/data"
	"errors"
)

func Logout() error {
	if data.Config.Basic.BaseUrl == "" {
		return errors.New("URL不能为空！")
	}
	return api.Logout(data.Config.Basic.BaseUrl + "/ac_portal/login.php")
}
