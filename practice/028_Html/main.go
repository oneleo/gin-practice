package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type myForm struct {
	Colors []string `form:"colors[]"`
}

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

	r.LoadHTMLGlob("./views/*")
	r.GET("/", indexHandler)
	r.POST("/", formHandler)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// 在第一次使用瀏覽器進入 http://localhost:8085 使會觸發 Get。
func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "form.html", nil)
}

// 點選 form.html 內 / 的 Post 鈕時，會觸發此函數。
func formHandler(c *gin.Context) {
	var fakeForm myForm
	c.Bind(&fakeForm)
	c.JSON(http.StatusOK, gin.H{"color": fakeForm.Colors})
}

// 使用瀏覽器瀏覽：
// http://localhost:8085
