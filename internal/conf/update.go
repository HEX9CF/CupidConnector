package conf

import (
	"CupidConnector/internal/model"
)

func Update(c model.Conf) error {
	Config = c
	return saveEnv()
}
