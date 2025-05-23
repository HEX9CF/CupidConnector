package api

import (
	"cupid-connector/backend/network"
	"cupid-connector/backend/parser"
	"cupid-connector/backend/utils"
	"errors"
	"log"
)

const LogoutOpr = "logout"

func Logout(url string) error {
	data := "opr=" + LogoutOpr
	bodyStr, err := network.PostRequest(url, data)
	if err != nil {
		return err
	}

	resp, err := parser.ParseLoginResp(bodyStr)
	if err != nil {
		return err
	}

	log.Println(utils.PrettyStruct(resp))

	if resp.Success != true {
		return errors.New("注销失败：" + resp.Msg)
	}

	log.Println("注销成功")
	return nil
}
