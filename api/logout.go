package api

import (
	"errors"
	"fmt"
	"log"
	"stu-campus-network-auto-login/network"
	"stu-campus-network-auto-login/utils"
)

const LogoutOpr = "logout"

func Logout(url string) error {
	data := "opr=" + LogoutOpr
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
		return errors.New("注销失败：" + resp.Msg)
	}

	log.Println("注销成功")
	return nil
}