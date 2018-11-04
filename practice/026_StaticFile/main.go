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
	r.Static("/assets", "./assets")
	r.StaticFS("/more_static", http.Dir("./my_file_system"))
	r.StaticFile("/favicon.ico", "./resources/favicon.ico")

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/assets"
/*
<a href="/assets/">Moved Permanently</a>.
*/

// $> curl -X GET "http://localhost:8085/more_static"
/*
<a href="/more_static/">Moved Permanently</a>.
*/

// $> curl -X GET "http://localhost:8085/favicon.ico"
/*
favicon.ico
*/
