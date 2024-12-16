package service

import (
	"cupid-connector/internal/conf"
	"cupid-connector/internal/model"
)

func UpdateConf(c model.Conf) error {
	return conf.Update(c)
}

func GetConf() model.Conf {
	return conf.Config
}
