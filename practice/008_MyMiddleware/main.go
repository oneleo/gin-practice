package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定義一個Looger中間件
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()

		// Set example variable
		c.Set("example", "12345")

		// request請求之前做什麼的代碼寫在這裡

		c.Next()

		// request請求之後做什麼的代碼寫在這裡
		latency := time.Since(t)
		log.Print(latency)

		// 獲取我們正在發送的狀態
		status := c.Writer.Status()
		log.Println(status)
	}
}

func main() {

	// set debug mode
	gin.SetMode(gin.DebugMode)

	// set release mode
	//gin.SetMode(gin.ReleaseMode)

	// 禁用控制台顏色，寫入日誌文件時不需要添加顏色
	// gin.DisableConsoleColor()

	r := gin.New()

	// Default With the Logger and Recovery middleware already attached
	// 默認情況已啟用了log和恢復中間件
	//r := gin.Default()

	// 默認情況下創建一個沒有任何中間件的路由器
	r.Use(Logger())

	r.GET("/test", func(c *gin.Context) {
		// 獲取中間件設置的變量
		example := c.MustGet("example").(string)

		// it would print: "12345"
		log.Println(example)
		c.String(http.StatusOK, example)
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/test"
/*
12345
*/
