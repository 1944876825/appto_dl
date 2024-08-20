package main

import (
	"appto_dl/config"
	"appto_dl/server"
	"appto_dl/utils"
	"fmt"
)

func main() {
	fmt.Println("教程：https://blog.prizen.cn/423.html")
	if utils.ExistFile("config.yaml") {
		config.Load()
	} else {
		utils.GetInfo()
	}
	server.Run()
}
