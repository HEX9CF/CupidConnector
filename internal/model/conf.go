package model

type Conf struct {
	Basic   BasicConf
	Monitor MonitorConf
}

type BasicConf struct {
	BaseUrl   string `json:"base_url"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin string `json:"auto_login"`
	AutoExit  string `json:"auto_exit"`
}

type MonitorConf struct {
	MonitorFlux     string `json:"monitor_flux"`
	MonitorInterval string `json:"monitor_interval"`
	AlertThreshold  string `json:"alert_threshold"`
	LogoutThreshold string `json:"logout_threshold"`
}
