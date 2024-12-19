package conf

import (
	"cupid-connector/internal/data"
	"os"
)

func getEnv() {
	data.Config.Basic.BaseUrl = os.Getenv("BASE_URL")
	data.Config.Basic.Username = os.Getenv("STU_USERNAME")
	data.Config.Basic.Password = os.Getenv("STU_PASSWORD")
	data.Config.Basic.AutoLogin = os.Getenv("AUTO_LOGIN")
	data.Config.Basic.AutoExit = os.Getenv("AUTO_EXIT")
	data.Config.Basic.AutoHide = os.Getenv("AUTO_HIDE")

	data.Config.Monitor.Enable = os.Getenv("MONITOR_FLUX")
	data.Config.Monitor.Interval = os.Getenv("MONITOR_INTERVAL")
	data.Config.Monitor.AlertThreshold = os.Getenv("MONITOR_ALERT_THRESHOLD")
	data.Config.Monitor.LogoutThreshold = os.Getenv("MONITOR_LOGOUT_THRESHOLD")
}

func getEnvContent() string {
	content := "BASE_URL=" + data.Config.Basic.BaseUrl + "\n" +
		"STU_USERNAME=" + data.Config.Basic.Username + "\n" +
		"STU_PASSWORD=" + data.Config.Basic.Password + "\n" +
		"AUTO_LOGIN=" + data.Config.Basic.AutoLogin + "\n" +
		"AUTO_EXIT=" + data.Config.Basic.AutoExit + "\n" +
		"MONITOR_FLUX=" + data.Config.Monitor.Enable + "\n" +
		"MONITOR_INTERVAL=" + data.Config.Monitor.Interval + "\n" +
		"MONITOR_ALERT_THRESHOLD=" + data.Config.Monitor.AlertThreshold + "\n" +
		"MONITOR_LOGOUT_THRESHOLD=" + data.Config.Monitor.LogoutThreshold + "\n" +
		"AUTO_HIDE=" + data.Config.Basic.AutoHide
	return content
}
