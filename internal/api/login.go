package api

import (
	"cupid-connector/internal/network"
	"cupid-connector/internal/utils"
	"errors"
	"fmt"
	"log"
)

const LoginOpr = "pwdLogin"

func Login(url string, username string, password string, rememberPwd string) error {
	data := "opr=" + LoginOpr + "&userName=" + username + "&pwd=" + password + "&rememberPwd=" + rememberPwd
	bodyStr, err := network.PostRequest(url, data)
	if err != nil {
		return err
	}

	resp, err := network.ResolveResponse(bodyStr)
	if err != nil {
		return err
	}

	fmt.Println(utils.PrettyStruct(resp))

	if resp.Success != true {
		return errors.New("登录失败：" + resp.Msg)
	}

	log.Println("登录成功，用户：" + username + "，您已通过上网认证！")
	return nil
}
