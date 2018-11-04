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

	r.POST("/form_post", func(c *gin.Context) {
		// 獲取 post 過來的 message 內容
		// 獲取的所有參數內容的類型都是 string
		message := c.PostForm("message")
		// 如果不存在，使用第二個當做默認內容
		nick := c.DefaultPostForm("nick", "anonymous")
		// 以 Json 輸出 Response。
		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"message": message,
			"nick":    nick,
		})
	})
	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X POST "http://localhost:8085/form_post" --form nick=bddbnet --form message=hello
// {"message":"hello","nick":"bddbnet","status":"posted"}

// $> curl -X POST "http://localhost:8085/form_post" -F message=hello
// {"message":"hello","nick":"anonymous","status":"posted"}

// $> curl -X POST "http://localhost:8085/form_post" -F nick=bddbnet
// {"message":"","nick":"bddbnet","status":"posted"}
