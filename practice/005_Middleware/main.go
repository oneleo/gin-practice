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

	// 註冊一個路由，使用了 middleware1，middleware2 兩個中間件
	r.GET("/someGet", middleware1, middleware2, handler)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

func handler(c *gin.Context) {
	log.Println("exec handler")
	c.String(http.StatusOK, "exec handler\n")
}

func middleware1(c *gin.Context) {
	log.Println("exec middleware 1")
	c.String(http.StatusOK, "exec middleware 1\n")
	//你可以寫一些邏輯代碼

	// 執行該中間件之後的邏輯
	c.Next()
}
func middleware2(c *gin.Context) {
	log.Println("arrive at middleware 2")
	c.String(http.StatusOK, "arrive at middleware 2\n")

	// 執行該中間件之前，先跳到流程的下一個方法
	c.Next()

	// 流程中的其他邏輯已經執行完了
	log.Println("exec middleware 2")
	c.String(http.StatusOK, "exec middleware 2\n")

	//你可以寫一些邏輯代碼
}

// $> curl -X GET "http://localhost:8085/someGet"
/*
exec middleware 1
arrive at middleware 2
exec handler
exec middleware 2
*/
