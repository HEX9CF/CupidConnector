package conf

import (
	"cupid-connector/backend/utils"
	"log"

	"github.com/joho/godotenv"
)

const envPath = "stu.env"

func InitEnv() error {
	log.Println("正在加载配置文件：" + envPath)
	v, err := utils.IsFileExists(envPath)
	if err != nil {
		log.Println("加载配置文件失败！" + err.Error())
		return err
	}

	if !v {
		log.Println("配置文件不存在！")
		log.Println("首次运行程序将创建配置文件")
		err := CreateEnv()
		if err != nil {
			return err
		}
	}

	err = godotenv.Load(envPath)
	if err != nil {
		log.Println("加载配置文件失败！")
		return err
	}

	getEnv()

	log.Println("加载配置文件成功")
	return nil
}
