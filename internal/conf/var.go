package conf

import (
	"cupid-connector/internal/model"
	"os"
)

var (
	Config model.Conf
)

func getEnv() {
	Config.Basic.BaseUrl = os.Getenv("BASE_URL")
	Config.Basic.Username = os.Getenv("STU_USERNAME")
	Config.Basic.Password = os.Getenv("STU_PASSWORD")
	Config.Basic.AutoLogin = os.Getenv("AUTO_LOGIN")
	Config.Basic.AutoExit = os.Getenv("AUTO_EXIT")

	Config.Monitor.Enable = os.Getenv("MONITOR_FLUX")
	Config.Monitor.Interval = os.Getenv("MONITOR_INTERVAL")
	Config.Monitor.AlertThreshold = os.Getenv("MONITOR_ALERT_THRESHOLD")
	Config.Monitor.LogoutThreshold = os.Getenv("MONITOR_LOGOUT_THRESHOLD")
}

func getEnvContent() string {
	content := "BASE_URL=" + Config.Basic.BaseUrl + "\n" +
		"STU_USERNAME=" + Config.Basic.Username + "\n" +
		"STU_PASSWORD=" + Config.Basic.Password + "\n" +
		"AUTO_LOGIN=" + Config.Basic.AutoLogin + "\n" +
		"AUTO_EXIT=" + Config.Basic.AutoExit + "\n" +
		"MONITOR_FLUX=" + Config.Monitor.Enable + "\n" +
		"MONITOR_INTERVAL=" + Config.Monitor.Interval + "\n" +
		"MONITOR_ALERT_THRESHOLD=" + Config.Monitor.AlertThreshold + "\n" +
		"MONITOR_LOGOUT_THRESHOLD=" + Config.Monitor.LogoutThreshold + "\n"
	return content
}
