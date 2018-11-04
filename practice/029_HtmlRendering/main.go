package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 用 LoadHTMLGlob() 或 LoadHTMLFiles() 函數加載模板文件

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

	// 加載所有的模板文件
	r.LoadHTMLGlob("templates/*")

	// 加載某個模板文件
	// router.LoadHTMLFiles("templates/template1.html", "templates/template2.html")

	r.GET("/index", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"title": "Main website", "context": "Web context",
		})
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// 使用瀏覽器瀏覽：
// http://localhost:8085/index
