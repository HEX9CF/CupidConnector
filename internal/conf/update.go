package conf

import (
	"cupid-connector/internal/model"
)

func Update(c model.Conf) error {
	Config = c
	return saveEnv()
}
