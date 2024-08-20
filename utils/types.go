package utils

type Json map[string]interface{}

type BaseRes struct {
	Status int `json:""`
	Msg    string
	Data   string
}

type AvailAbilityRes struct {
	Data struct {
		Features   []string `json:"features"`
		MacAddress string   `json:"mac_address"`
		Type       int      `json:"type"`
	} `json:"data"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
