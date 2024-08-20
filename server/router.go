package server

import (
	"appto_dl/config"
	"fmt"

	"github.com/gin-gonic/gin"
)

func Run() {
	gin.SetMode(gin.ReleaseMode)
	r := gin.Default()

	r.Any("/ping", func(c *gin.Context) {
		c.String(200, "pong")
	})
	r.GET("/addon/v1/ability", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "OK",
			"data":   config.Conf.Endata,
		})
	})
	r.PUT("/addon/v1/mac_address", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "OK",
			"data":   nil,
		})
	})
	r.GET("/addon/v1/availability", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"status": 0,
			"msg":    "OK",
			"data": gin.H{
				"features": []string{
					"epay",
					"dmku",
					"see_ad_award",
					"cloud_history",
					"casd_sdk",
					"promotion",
					"page_favorite_custom_1",
				},
				"mac_address": config.Conf.EnMac,
				"type":        2,
			},
		})
	})

	// 自定义 404 错误处理函数
	r.NoRoute(func(c *gin.Context) {
		// 获取请求的方法
		method := c.Request.Method

		// 获取访问的路由
		path := c.Request.URL.Path

		// 获取查询参数
		queryParams := c.Request.URL.Query()

		// 获取表单参数（针对 POST, PUT 等请求）
		formParams := c.Request.PostForm

		// 尝试解析 JSON 数据
		var jsonBody map[string]interface{}
		if err := c.ShouldBindJSON(&jsonBody); err != nil {
			jsonBody = nil // 如果 JSON 解析失败，置为 nil
		}

		// 获取请求的协议头
		headers := c.Request.Header

		// 打印访问的路由、参数、协议头等信息
		fmt.Printf("Method: %s\n", method)
		fmt.Printf("Path: %s\n", path)
		fmt.Printf("Query Params: %v\n", queryParams)
		fmt.Printf("Form Params: %v\n", formParams)
		fmt.Printf("JSON Body: %v\n", jsonBody)
		fmt.Printf("Headers: %v\n", headers)

		// 返回 JSON 响应
		c.JSON(200, gin.H{
			"status": 404,
			"msg":    "not router",
		})
	})
	fmt.Println("程序启动成功: ", fmt.Sprintf("http://127.0.0.1:%d", config.Conf.Port))
	err := r.Run(fmt.Sprintf(":%d", config.Conf.Port))
	if err != nil {
		panic("web服务启动失败，" + err.Error())
	}
}

// 获取绑定域名列表
// get /addon/v1/certificate/domains
// {"status":0,"msg":"OK","data":{"domains":["http://v.qq.com"]}}

// 添加绑定域名
// put /addon/v1/certificate/domains
// {
// "domains":["http://v.qq.com"]
// }
// {"status":0,"msg":"OK","data":null}
