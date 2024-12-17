package conf

const (
	defaultBaseUrl         = "https://a.stu.edu.cn"
	defaultUsername        = ""
	defaultPassword        = ""
	defaultAutoLogin       = "TRUE"
	defaultAutoExit        = "FALSE"
	defaultMonitorFlux     = "FALSE"
	defaultMonitorInterval = "5"
	defaultAlertThreshold  = "1024"
	defaultLogoutInterval  = "512"
)

func setDefault() {
	Config.Basic.BaseUrl = defaultBaseUrl
	Config.Basic.Username = defaultUsername
	Config.Basic.Password = defaultPassword
	Config.Basic.AutoLogin = defaultAutoLogin
	Config.Basic.AutoExit = defaultAutoExit
	Config.Monitor.MonitorFlux = defaultMonitorFlux
	Config.Monitor.MonitorInterval = defaultMonitorInterval
	Config.Monitor.AlertThreshold = defaultAlertThreshold
	Config.Monitor.LogoutThreshold = defaultLogoutInterval
}
