package utils

type Json map[string]interface{}

type BaseRes struct {
	Status int    `json:"status"`
	Msg    string `json:"msg"`
	Data   string `json:"data"`
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

type DomainsRes struct {
	Data struct {
		Domains []string `json:"domains"`
	} `json:"data"`
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}
