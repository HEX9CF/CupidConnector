package service

import (
	"CupidConnector/internal/conf"
	"CupidConnector/internal/model"
)

func UpdateConf(c model.Conf) error {
	return conf.Update(c)
}

func GetConf() model.Conf {
	return conf.Config
}
