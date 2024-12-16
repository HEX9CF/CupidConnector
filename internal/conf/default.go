package conf

const (
	defaultUrl      = "http://a.stu.edu.cn/ac_portal/login.php"
	defaultUsername = "username"
	defaultPasswrod = "password"
	defaultAutoExit = "FALSE"
)

func setDefault() {
	Config.Url = defaultUrl
	Config.Username = defaultUsername
	Config.Password = defaultPasswrod
	Config.AutoExit = defaultAutoExit
}
