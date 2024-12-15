package conf

const (
	defaultUrl      = "http://a.stu.edu.cn/ac_portal/login.php"
	defaultUsername = "username"
	defaultPasswrod = "password"
	defaultAutoExit = "FALSE"
)

func setDefault() {
	Url = defaultUrl
	Username = defaultUsername
	Password = defaultPasswrod
	AutoExit = defaultAutoExit
}
