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

	// 註冊路由和Handler
	// url 為 /welcome?firstname=Jane&lastname=Doe
	r.GET("/welcome", func(c *gin.Context) {
		// 獲取參數內容
		// 獲取的所有參數內容的類型都是 string
		// 如果不存在，使用第二個當做默認內容
		firstname := c.DefaultQuery("firstname", "Guest")
		// 獲取參數內容，沒有則返回空字符串
		lastname := c.Query("lastname")
		// 以 String 輸出 Response。
		c.String(http.StatusOK, "Hello, %s %s\n", firstname, lastname)
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/welcome?firstname=Jane&lastname=Doe"
// $> curl -X GET "http://localhost:8085/welcome?&lastname=Doe"
// $> curl -X GET "http://localhost:8085/welcome"
