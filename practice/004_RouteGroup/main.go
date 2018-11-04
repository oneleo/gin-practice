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

	// 定義一個組前綴
	// /v1/login 就會匹配到這個組
	v1 := r.Group("/v1")
	{
		v1.POST("/login", loginEndpoint)
		v1.POST("/submit", submitEndpoint)
		v1.POST("/read", readEndpoint)
	}

	// 定義一個組前綴
	// 不用花括號包起來也是可以的。上面那種只是看起來會統一一點。看你個人喜好
	v2 := r.Group("/v2")
	v2.POST("/login", loginEndpoint)
	v2.POST("/submit", submitEndpoint)
	v2.POST("/read", readEndpoint)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
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

// $> curl -X POST "http://localhost:8085/v1/login"
// $> curl -X POST "http://localhost:8085/v1/submit"
// $> curl -X POST "http://localhost:8085/v1/read"

// $> curl -X POST "http://localhost:8085/v2/login"
// $> curl -X POST "http://localhost:8085/v2/submit"
// $> curl -X POST "http://localhost:8085/v2/read"
