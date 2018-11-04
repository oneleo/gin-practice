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

	r.GET("/someGet", getting)
	r.POST("/somePost", posting)
	r.PUT("/somePut", putting)
	r.DELETE("/someDelete", deleting)
	r.PATCH("/somePatch", patching)
	r.HEAD("/someHead", head)
	r.OPTIONS("/someOptions", options)

	// By default it serves on :8080 unless a
	// PORT environment variable was defined.
	// 服務默認使用 8080 端口，除非你自定義了端口號的環境變量
	//r.Run()

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

func getting(c *gin.Context) {
	c.String(http.StatusOK, "getting")
}
func posting(c *gin.Context) {
	c.String(http.StatusOK, "posting")
}
func putting(c *gin.Context) {
	c.String(http.StatusOK, "putting")
}
func deleting(c *gin.Context) {
	c.String(http.StatusOK, "deleting")
}
func patching(c *gin.Context) {
	c.String(http.StatusOK, "patching")
}
func head(c *gin.Context) {
	c.String(http.StatusOK, "head")
}
func options(c *gin.Context) {
	c.String(http.StatusOK, "options")
}

// $> curl -X GET "http://localhost:8085/someGet"
// $> curl -X POST "http://localhost:8085/somePost"
// $> curl -X PUT "http://localhost:8085/somePut"
// $> curl -X DELETE "http://localhost:8085/someDelete"
// $> curl -X PATCH "http://localhost:8085/somePatch"
// $> curl -X OPTIONS "http://localhost:8085/someOptions"

// $> curl -X HEAD --head "http://localhost:8085/someHead"
/*
HTTP/1.1 200 OK
Content-Type: text/plain; charset=utf-8
Date: Sat, 03 Nov 2018 05:50:53 GMT
Content-Length: 4
*/
