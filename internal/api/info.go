package api

import (
	"cupid-connector/internal/model"
	"cupid-connector/internal/network"
	"cupid-connector/internal/parser"
	"cupid-connector/internal/utils"
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
