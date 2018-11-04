package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {

	// set debug mode
	gin.SetMode(gin.DebugMode)

	// set release mode
	//gin.SetMode(gin.ReleaseMode)

	// 禁用控制台顏色，寫入日誌文件時不需要添加顏色
	// gin.DisableConsoleColor()

	// 初始化引擎。
	// 用默認的中間件創建一個gin路由器:
	// logger and recovery (crash-free) middleware
	// 記錄 恢復(不崩潰) 中間件
	r := gin.Default()

	// 註冊一個動態路由
	// 路由 1：匹配「/user/john」，但是不匹配「/user/」或「 /user」
	// 注意：「/user/:name」和「/user/:name/」是倆個完全不同的路由
	r.GET("/user/:name", func(c *gin.Context) {
		// 使用 c.Param(key) 獲取 url 參數
		name := c.Param("name")
		// 以 String 輸出 Response。
		c.String(http.StatusOK, "Hello %s", name)
	})

	// 註冊一個高級的動態路由
	// 路由 2 該路由會匹配 /user/john/ 和 /user/john/send
	// 如果沒有任何路由匹配到 /user/john, 那麼他就會重定向到 /user/john/，從而被該方法匹配到
	r.GET("/user/:name/*action", func(c *gin.Context) {
		name := c.Param("name")
		action := c.Param("action")
		message := name + " is " + action
		// 以 String 輸出 Response。
		c.String(http.StatusOK, message)
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/user/jack"
// Hello jack

// $> curl -X GET "http://localhost:8085/user/jack/"
// jack is /

// $> curl -X GET "http://localhost:8085/user/jack/do"
// jack is /do
