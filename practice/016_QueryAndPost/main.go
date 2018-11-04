package main

import (
	"fmt"
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

	r.POST("/post", func(c *gin.Context) {
		// url中查詢數據
		id := c.Query("id")
		page := c.DefaultQuery("page", "0")

		// post表單中數據
		name := c.PostForm("name")
		message := c.PostForm("message")

		c.JSON(http.StatusOK, gin.H{
			"status":  "posted",
			"id":      id,
			"page":    page,
			"name":    name,
			"message": message,
		})
		fmt.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
		log.Printf("id: %s; page: %s; name: %s; message: %s\n", id, page, name, message)
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X POST "http://localhost:8085/post?id=123&page=1" -F name=bddbnet -F message=hello
// {"id":"123","message":"hello","name":"bddbnet","page":"1","status":"posted"}

// $> curl -X POST "http://localhost:8085/post?id=123" -F name=bddbnet -F message=hello
//{"id":"123","message":"hello","name":"bddbnet","page":"0","status":"posted"}

// $> curl -X POST "http://localhost:8085/post" -F name=bddbnet -F message=hello
//{"id":"","message":"hello","name":"bddbnet","page":"0","status":"posted"}

// $> curl -X POST "http://localhost:8085/post" -F name=bddbnet
//{"id":"","message":"","name":"bddbnet","page":"0","status":"posted"}
