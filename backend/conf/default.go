package conf

import "cupid-connector/backend/data"

const (
	DefaultBaseUrl  = "https://a.stu.edu.cn"
	DefaultUsername = ""
	DefaultPassword = ""

	DefaultAutoLogin     = "FALSE"
	DefaultAutoExit      = "FALSE"
	DefaultAutoHide      = "FALSE"
	DefaultAutoReconnect = "FALSE"

	DefaultMonitorFlux     = "FALSE"
	DefaultMonitorInterval = "5"
	DefaultAlertThreshold  = "30"
	DefaultLogoutThreshold = "10"
)

func setDefault() {
	data.Config.Basic.BaseUrl = DefaultBaseUrl
	data.Config.Basic.Username = DefaultUsername
	data.Config.Basic.Password = DefaultPassword
	data.Config.Basic.AutoLogin = DefaultAutoLogin
	data.Config.Basic.AutoExit = DefaultAutoExit
	data.Config.Basic.AutoHide = DefaultAutoHide
	data.Config.Monitor.AutoReconnect = DefaultAutoReconnect

	data.Config.Monitor.Enable = DefaultMonitorFlux
	data.Config.Monitor.Interval = DefaultMonitorInterval
	data.Config.Monitor.AlertThreshold = DefaultAlertThreshold
	data.Config.Monitor.LogoutThreshold = DefaultLogoutThreshold
}
