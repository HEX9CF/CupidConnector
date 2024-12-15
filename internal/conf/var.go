package conf

import "os"

var (
	Url      string
	Username string
	Password string
	AutoExit string
)

func getEnv() {
	Url = os.Getenv("STU_URL")
	Username = os.Getenv("STU_USERNAME")
	Password = os.Getenv("STU_PASSWORD")
	AutoExit = os.Getenv("AUTO_EXIT")
}
