package conf

const (
	defaultUrl      = "http://a.stu.edu.cn/ac_portal/login.php"
	defaultUsername = ""
	defaultPasswrod = ""
	defaultAutoExit = "FALSE"
)

func setDefault() {
	Config.Url = defaultUrl
	Config.Username = defaultUsername
	Config.Password = defaultPasswrod
	Config.AutoExit = defaultAutoExit
}
