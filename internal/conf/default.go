package conf

const (
	defaultBaseUrl   = "https://a.stu.edu.cn"
	defaultUsername  = ""
	defaultPassword  = ""
	defaultAutoLogin = "TRUE"
	defaultAutoExit  = "FALSE"
)

func setDefault() {
	Config.BaseUrl = defaultBaseUrl
	Config.Username = defaultUsername
	Config.Password = defaultPassword
	Config.AutoLogin = defaultAutoLogin
	Config.AutoExit = defaultAutoExit
}
