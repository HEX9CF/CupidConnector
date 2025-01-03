package api

import (
	"cupid-connector/backend/network"
	"cupid-connector/backend/parser"
	"cupid-connector/backend/utils"
	"errors"
	"log"
)

const LoginOpr = "pwdLogin"

func Login(url string, username string, password string, rememberPwd string) error {
	data := "opr=" + LoginOpr + "&userName=" + username + "&pwd=" + password + "&rememberPwd=" + rememberPwd
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
		return errors.New("登录失败，" + resp.Msg)
	}

	log.Println("登录成功，用户：" + resp.UserName + "，您已通过上网认证！")
	return nil
}
