package utils

import (
	"appto_dl/config"
	"bytes"
	"crypto/md5"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net"
	"net/http"
	"net/url"
	"strings"
)

func GetInfo() {
	getInput()
	getAvailAbility()
	fmt.Println("云端MacAddress是:", config.Conf.EnMac)
	macAddress, err := getMACAddress()
	if err != nil {
		panic("获取本地mac地址失败" + err.Error())
	}
	fmt.Println("本机MacAddress是:", macAddress)
	fmt.Printf("是否设置Mac地址（y/n）：")
	var zl string
	if _, err := fmt.Scanln(&zl); err != nil {
		panic(err)
	}
	if strings.ToLower(zl) == "y" {
		fmt.Printf("请输入你要设置的MacAddress，32位：")
		var zdyMac string
		if _, err := fmt.Scanln(&zdyMac); err != nil {
			panic(err)
		}
		zdyMac = strings.TrimSpace(zdyMac)
		if zdyMac == "" {
			zdyMac = macAddress
		}
		setMacAddress(zdyMac)
		fmt.Println("设置成功")
	}
	getAbility()
	GetDomains()
	if len(config.Conf.Domains) < 1 {
		fmt.Println("未获取到你绑定的域名")
	} else {
		fmt.Println("你当前绑定了以下域名")
		for _, v := range config.Conf.Domains {
			fmt.Println(v)
		}
	}
	fmt.Printf("是否重新绑定（y/n）:")
	if _, err := fmt.Scanln(&zl); err != nil {
		panic(err)
	}
	if strings.ToLower(zl) == "y" {
		fmt.Println("请输入需要绑定的域名，#隔开")
		var dm string
		if _, err := fmt.Scanln(&dm); err != nil {
			panic(err)
		}
		dms := strings.Split(dm, "#")
		fmt.Printf("检测到%d个域名\n", len(dms))
		setDomains(dms)
		fmt.Println("设置成功")
	}
	config.Save()
}
func getInput() {
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
	base             = "http://v3.appto.top"
	abilityApi       = base + "/addon/v1/ability"
	availabilityApi  = base + "/addon/v1/availability"
	setMacAddressApi = base + "/addon/v1/mac_address"
	getDomainsApi    = base + "/addon/v1/certificate/domains"
	setDomainsApi    = base + "/addon/v1/certificate/domains"
)

// 获取Ability 其中的data是一个加密信息，可能包含token、mac地址、插件授权信息等。
func getAbility() {
	res, err := request(abilityApi, "GET", nil)
	if err != nil {
		panic(err)
	}
	var r BaseRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		log.Fatalln(err, string(res))
	}
	if r.Status != 0 {
		panic(r)
	}
	config.Conf.Endata = r.Data
}

// 获取 AvailAbility
func getAvailAbility() {
	res, err := request(availabilityApi, "GET", nil)
	if err != nil {
		panic(err)
	}
	var r AvailAbilityRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		log.Fatalln(err, string(res))
	}
	if r.Status != 0 {
		panic(r)
	}
	config.Conf.EnMac = r.Data.MacAddress
}

// 设置MacAddress
func setMacAddress(macAddress string) {
	data := Json{
		"mac_address": macAddress,
	}
	res, err := request(setMacAddressApi, "PUT", data)
	if err != nil {
		panic(err)
	}
	// fmt.Println("s", string(res))
	var r BaseRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		log.Fatalln(err, string(res))
	}
	if r.Status != 0 {
		panic(r)
	}
}

// 设置绑定域名
func setDomains(dms []string) {
	data := Json{
		"domains": dms,
	}
	res, err := request(setDomainsApi, "PUT", data)
	if err != nil {
		panic(err)
	}
	var r BaseRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		log.Fatalln(err, string(res))
	}
	if r.Status != 0 {
		panic(r)
	}
}

// getMACAddress 尝试获取本机的MAC地址
func getMACAddress() (string, error) {
	ifaces, err := net.Interfaces()
	if err != nil {
		return "", err
	}
	for _, iface := range ifaces {
		if iface.Flags&net.FlagUp != 0 && iface.Flags&net.FlagLoopback == 0 {
			return md5Encode(iface.HardwareAddr.String()), nil
		}
	}
	return "", fmt.Errorf("获取不到Mac地址")
}

// 获取域名列表
func GetDomains() {
	res, err := request(getDomainsApi, "GET", nil)
	if err != nil {
		panic(err)
	}
	var r DomainsRes
	err = json.Unmarshal(res, &r)
	if err != nil {
		log.Fatalln(err, string(res))
	}
	if r.Status != 0 {
		panic(r)
	}
	config.Conf.Domains = r.Data.Domains
}

var ProxyUrl *url.URL

// setProxy 设置http代理
func setProxy(p string) {
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
	req.Header.Set("Content-Type", "application/json")
	// setProxy("127.0.0.1:7890")
	// transport := &http.Transport{
	// 	Proxy: http.ProxyURL(ProxyUrl),
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
func md5Encode(text string) string {
	hash := md5.New()
	hash.Write([]byte(text))
	return fmt.Sprintf("%x", hash.Sum(nil))
}
