package model

type Info struct {
	Username       string  `json:"user_name"`
	Overall        float64 `json:"overall"`
	Used           float64 `json:"used"`
	ExpirationTime string  `json:"expiration_time"`
	AccountStatus  string  `json:"account_status"`
}
