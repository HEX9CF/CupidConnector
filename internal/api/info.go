package api

import (
	"cupid-connector/internal/model"
	"cupid-connector/internal/network"
	"cupid-connector/internal/parser"
)

func GetInfo() (model.Info, error) {
	body, err := network.PostRequest("https://a.stu.edu.cn/ac_portal/userflux", "")
	if err != nil {
		return model.Info{}, err
	}
	info := parser.ParseFluxInfo(body)
	return info, nil
}
