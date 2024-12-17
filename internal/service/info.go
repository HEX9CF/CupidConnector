package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/conf"
	"cupid-connector/internal/model"
)

func GetInfo() (model.Info, error) {
	return api.GetInfo(conf.Config.BaseUrl + "/ac_portal/userflux")
}
