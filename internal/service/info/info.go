package info

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/data"
)

func Refresh() error {
	var err error
	data.Info, err = api.GetInfo(data.Config.Basic.BaseUrl + "/ac_portal/userflux")
	if err != nil {
		return err
	}
	return nil
}
