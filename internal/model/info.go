package model

type Info struct {
	Username       string `json:"user_name"`
	Overall        string `json:"overall"`
	Used           string `json:"used"`
	ExpirationTime string `json:"expiration_time"`
	AccountStatus  string `json:"account_status"`
}
