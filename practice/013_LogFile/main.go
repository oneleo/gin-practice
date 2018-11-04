package main

import (
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func main() {

	// 禁用控制台顏色，寫入日誌文件時不需要添加顏色
	gin.DisableConsoleColor()

	// 寫入到文件。
	f, _ := os.Create("./gin.log")
	gin.DefaultWriter = io.MultiWriter(f)

	// 如果您需要同時將日誌寫入文件和控制台，請使用以下代碼
	// gin.DefaultWriter = io.MultiWriter(f, os.Stdout)

	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "ping")
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/ping"
/*
exec MyBenchLogger
Hello
*/
