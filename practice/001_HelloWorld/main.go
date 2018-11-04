package main

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"

	"net/http"
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
	router := gin.Default()

	// 註冊一個路由和處理函數。
	irRoot := router.Any("/", WebRoot)
	fmt.Println("Web: / Info:\n", irRoot)

	irData := router.GET("/ping", WebPing)
	fmt.Println("Web: /ping Info:\n", irData)

	// 綁定端口，然後啟動應用。
	err := router.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// 根 / 請求處理函數
// 所有本次請求相關的方法都在 context 中，
// 完美輸出響應 hello, world。

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}
func WebPing(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "ping",
	})
}

// $> curl --request GET "http://localhost:8085"
// hello, world

// $> curl -X GET "http://localhost:8085/ping"
// {"message":"ping"}
