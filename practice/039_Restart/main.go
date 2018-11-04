package main

import (
	"fmt"
	"log"

	"github.com/facebookgo/grace/gracehttp"
	"github.com/gin-gonic/gin"

	"net/http"
)

func main() {

	// set debug mode
	gin.SetMode(gin.DebugMode)

	// 初始化引擎。
	router := gin.Default()

	// 註冊一個路由和處理函數。
	irRoot := router.Any("/", WebRoot)
	fmt.Println("Web: / Info:\n", irRoot)

	// 綁定端口，然後啟動應用。
	err := router.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}

	// 優雅的重啟服務
	// 注意：過程使用到 Linux 的 SIGNAL，並不支援 Windows 作業系統。
	gracehttp.Serve(
		&http.Server{Addr: ":8085", Handler: router},
	)

	// endless.ListenAndServe(":8085", router)
}

// 根 / 請求處理函數
// 所有本次請求相關的方法都在 context 中，
// 完美輸出響應 hello, world。

func WebRoot(context *gin.Context) {
	context.String(http.StatusOK, "hello, world")
}

// Needed:
// $> go get -u -v github.com/facebookgo/grace
// $> go get -u -v github.com/facebookgo/httpdown
