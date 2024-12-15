package model

type Response struct {
	Success  bool   `json:"success,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Action   string `json:"action,omitempty"`
	Pop      int    `json:"pop,omitempty"`
	UserName string `json:"userName,omitempty"`
	Location string `json:"location,omitempty"`
}
