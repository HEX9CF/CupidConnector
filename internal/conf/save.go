package conf

import (
	"cupid-connector/internal/utils"
	"log"
	"os"
)

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

	content := getEnvContent()

	_, err = file.WriteString(content)
	if err != nil {
		log.Println("写入配置文件失败！")
		return err
	}

	log.Println("创建配置文件成功！")
	return nil
}
