package service

import (
	"cupid-connector/internal/api"
	"cupid-connector/internal/model"
)

func GetInfo() (model.Info, error) {
	return api.GetInfo()
}
