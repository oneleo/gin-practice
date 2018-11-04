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

	// 定義一個組前綴, 並使用 middleware1 中間件
	// 訪問 /v2/login 就會執行 middleware1 函數
	v2 := r.Group("/v2", middleware1)
	v2.POST("/login", loginEndpoint)
	v2.POST("/submit", submitEndpoint)
	v2.POST("/read", readEndpoint)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}
func middleware1(c *gin.Context) {
	log.Println("exec middleware 1")
	c.String(http.StatusOK, "exec middleware 1\n")
	//你可以寫一些邏輯代碼

	// 執行該中間件之後的邏輯
	c.Next()
}
func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}
func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}
func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}

// $> curl -X POST "http://localhost:8085/v2/login"
// $> curl -X POST "http://localhost:8085/v2/submit"
// $> curl -X POST "http://localhost:8085/v2/read"
/*
exec middleware 1
URL: /v2/read
*/
