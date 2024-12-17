package conf

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
)

func setDefault() {
	Config.Basic.BaseUrl = DefaultBaseUrl
	Config.Basic.Username = DefaultUsername
	Config.Basic.Password = DefaultPassword
	Config.Basic.AutoLogin = DefaultAutoLogin
	Config.Basic.AutoExit = DefaultAutoExit
	Config.Monitor.Enable = DefaultMonitorFlux
	Config.Monitor.Interval = DefaultMonitorInterval
	Config.Monitor.AlertThreshold = DefaultAlertThreshold
	Config.Monitor.LogoutThreshold = DefaultLogoutThreshold
}
