package conf

import (
	"CupidConnector/internal/utils"
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
)

func CreateEnv() error {
	reader := bufio.NewReader(os.Stdin)
	for reader.Buffered() > 0 {
		_, err := reader.Discard(reader.Buffered())
		if err != nil {
			log.Println(err)
			break
		}
	}

	setDefault()

	fmt.Println("请输入校园网认证信息：")
	fmt.Print("用户名：")
	Username, _ = reader.ReadString('\n')
	Username = strings.TrimSpace(Username)
	fmt.Print("密码：")
	Password, _ = reader.ReadString('\n')
	Password = strings.TrimSpace(Password)

	AutoExit = defaultAutoExit
	fmt.Print("认证成功后是否自动退出程序(Y/N)？")
	autoQuit, _ := reader.ReadString('\n')
	autoQuit = strings.TrimSpace(autoQuit)
	autoQuit = strings.ToLower(autoQuit)
	if autoQuit == "y" {
		AutoExit = "TRUE"
	}

	err := saveEnv()
	if err != nil {
		return err
	}

	return nil
}

func saveEnv() error {
	log.Println("正在创建配置文件...")
	v, err := utils.IsFileExists(envPath)
	if err != nil {
		log.Println("创建配置文件失败！")
		return err
	}
	if v {
		err := os.Remove(envPath)
		if err != nil {
			log.Println("删除旧配置文件失败！")
			return err
		}
	}
	file, err := os.Create(envPath)
	if err != nil {
		log.Println("创建配置文件失败！")
		return err
	}
	defer file.Close()

	content := "STU_URL=" + defaultUrl + "\n" +
		"STU_USERNAME=" + Username + "\n" +
		"STU_PASSWORD=" + Password + "\n" +
		"AUTO_EXIT=" + AutoExit + "\n"

	_, err = file.WriteString(content)
	if err != nil {
		log.Println("写入配置文件失败！")
		return err
	}

	log.Println("创建配置文件成功！")
	return nil
}
