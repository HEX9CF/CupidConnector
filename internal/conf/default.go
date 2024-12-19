package conf

import "cupid-connector/internal/data"

const (
	DefaultBaseUrl         = "https://a.stu.edu.cn"
	DefaultUsername        = ""
	DefaultPassword        = ""
	DefaultAutoLogin       = "TRUE"
	DefaultAutoExit        = "FALSE"
	DefaultMonitorFlux     = "FALSE"
	DefaultMonitorInterval = "5"
	DefaultAlertThreshold  = "30"
	DefaultLogoutThreshold = "10"
	DefaultAutoHide        = "FALSE"
)

func setDefault() {
	data.Config.Basic.BaseUrl = DefaultBaseUrl
	data.Config.Basic.Username = DefaultUsername
	data.Config.Basic.Password = DefaultPassword
	data.Config.Basic.AutoLogin = DefaultAutoLogin
	data.Config.Basic.AutoExit = DefaultAutoExit
	data.Config.Basic.AutoHide = DefaultAutoHide
	data.Config.Monitor.Enable = DefaultMonitorFlux
	data.Config.Monitor.Interval = DefaultMonitorInterval
	data.Config.Monitor.AlertThreshold = DefaultAlertThreshold
	data.Config.Monitor.LogoutThreshold = DefaultLogoutThreshold
}
