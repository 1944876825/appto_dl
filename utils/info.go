package utils

import (
	"appto_dl/config"
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
	"strings"
)

func GetInfo() {
	getConfig()
	getAbility()
	getAvailAbility()
	config.Save()
}
func getConfig() {
	fmt.Printf("请输入运行的端口：")
	_, err := fmt.Scanln(&config.Conf.Port)
	if err != nil {
		panic(err)
	}
	fmt.Printf("请输入你的Token：")
	_, err = fmt.Scanln(&config.Conf.Token)
	if err != nil {
		panic(err)
	}
}

const (
	base            = "http://v3.appto.top"
	abilityApi      = base + "/addon/v1/ability"
	availabilityApi = base + "/addon/v1/availability"
)

func getAbility() {
	res, err := request(abilityApi, "GET", nil)
	if err != nil {
		panic(err)
	}
	var r BaseRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		panic(err)
	}
	if r.Status != 0 {
		panic(r)
	}
	config.Conf.Endata = r.Data
}
func getAvailAbility() {
	res, err := request(availabilityApi, "GET", nil)
	if err != nil {
		panic(err)
	}
	var r AvailAbilityRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		panic(err)
	}
	if r.Status != 0 {
		panic(r)
	}
	config.Conf.EnMac = r.Data.MacAddress
}

func GetDomains() {
	res, err := request(availabilityApi, "GET", nil)
	if err != nil {
		panic(err)
	}
	var r AvailAbilityRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		panic(err)
	}
	if r.Status != 0 {
		panic(r)
	}
	config.Conf.EnMac = r.Data.MacAddress
}

var ProxyUrl *url.URL

func SetProxy(p string) {
	var err error
	if strings.Contains(p, "http") == false {
		p = "http://" + p
	}
	ProxyUrl, err = url.Parse(p)
	if err != nil {
		log.Println("代理设置失败", err.Error())
	}
}
func request(url, method string, data interface{}) ([]byte, error) {
	var body io.Reader
	if data == nil {
		body = nil
	} else if val, ok1 := data.(string); ok1 {
		body = strings.NewReader(val)
	} else if val, ok2 := data.(Json); ok2 {
		marshal, err := json.Marshal(val)
		if err != nil {
			return nil, err
		}
		body = bytes.NewReader(marshal)
	}
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Access-Token", config.Conf.Token)
	// SetProxy("127.0.0.1:7890")
	// transport := &http.Transport{
	// Proxy: http.ProxyURL(ProxyUrl),
	// }
	client := http.Client{
		// Transport: transport,
	}
	do, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer do.Body.Close()
	res, err := io.ReadAll(do.Body)
	// log.Println("res1", string(res))
	if err != nil {
		return nil, err
	}
	return res, nil
}
