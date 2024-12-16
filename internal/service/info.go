package service

import (
	"CupidConnector/internal/api"
	"CupidConnector/internal/model"
)

func GetInfo() (model.Info, error) {
	return api.GetInfo()
}
