package model

type Conf struct {
	Url      string `json:"url"`
	Username string `json:"username"`
	Password string `json:"password"`
	AutoExit string `json:"auto_exit"`
}
