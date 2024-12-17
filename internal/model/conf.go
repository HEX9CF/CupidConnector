package model

type Conf struct {
	BaseUrl   string `json:"base_url"`
	Username  string `json:"username"`
	Password  string `json:"password"`
	AutoLogin string `json:"auto_login"`
	AutoExit  string `json:"auto_exit"`
}
