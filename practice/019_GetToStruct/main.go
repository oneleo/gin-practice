package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定義一個 Person 結構體，用來綁定 url query
type Person struct {
	Name    string `form:"name"` // 使用成員變量標籤定義對應的參數名
	Address string `form:"address"`
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
	r.Any("/testing", startPage)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

func startPage(c *gin.Context) {
	var person Person
	// 將 url 查詢參數和person綁定在一起
	if c.ShouldBindQuery(&person) == nil {
		log.Println("====== Only Bind By Query String ======")
		log.Println(person.Name)
		log.Println(person.Address)
	}
	// 以 String 回傳 Response。
	c.String(http.StatusOK, "Success\n")
	c.JSON(http.StatusOK, gin.H{"Name": person.Name, "Address": person.Address})
}

// $> curl -X GET "http://localhost:8085/testing?name=Irara&address=Taipei"
/*
{"Address":"Taipei","Name":"Irara"}
*/
