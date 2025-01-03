package api

import (
	"cupid-connector/backend/model"
	"cupid-connector/backend/network"
	"cupid-connector/backend/parser"
	"cupid-connector/backend/utils"
	"log"
)

func GetInfo(url string) (model.Info, error) {
	body, err := network.PostRequest(url, "")
	if err != nil {
		return model.Info{}, err
	}
	info := parser.ParseFluxInfo(body)

	log.Print(utils.PrettyStruct(info))
	return info, nil
}
