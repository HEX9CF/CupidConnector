package model

type ResponseCode uint8

const (
	ResponseCodeError ResponseCode = 0
	ResponseCodeOk    ResponseCode = 1
	ResponseCodeRetry ResponseCode = 2
)

func (c ResponseCode) String() string {
	switch c {
	case ResponseCodeError:
		return "错误"
	case ResponseCodeOk:
		return "成功"
	case ResponseCodeRetry:
		return "重新请求"
	default:
		return "未知状态"
	}
}

type Response struct {
	Success  bool   `json:"success,omitempty"`
	Msg      string `json:"msg,omitempty"`
	Action   string `json:"action,omitempty"`
	Pop      int    `json:"pop,omitempty"`
	UserName string `json:"userName,omitempty"`
	Location string `json:"location,omitempty"`
}

type Resp struct {
	Code ResponseCode `json:"code"`
	Msg  string       `json:"msg"`
	Data interface{}  `json:"data"`
}
