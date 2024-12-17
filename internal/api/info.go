package api

import (
	"cupid-connector/internal/model"
	"cupid-connector/internal/network"
	"cupid-connector/internal/parser"
)

func GetInfo(url string) (model.Info, error) {
	body, err := network.PostRequest(url, "")
	if err != nil {
		return model.Info{}, err
	}
	info := parser.ParseFluxInfo(body)
	return info, nil
}
