package main

import (
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 定義一個 Person 結構體，用來綁定數據
type Person struct {
	Name     string    `form:"name"`
	Address  string    `form:"address"`
	Birthday time.Time `form:"birthday" time_format:"2006-01-02" time_utc:"1"`
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

	// If `GET`, only `Form` binding engine (`query`) used.
	// If `POST`, first checks the `content-type` for `JSON` or `XML`, then uses `Form` (`form-data`).
	// See more at https://github.com/gin-gonic/gin/blob/master/binding/binding.go#L48
	r.GET("/testGet", startPage)
	r.POST("/testPost", startPage)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

func startPage(c *gin.Context) {
	var person Person
	// 綁定到 person
	if c.ShouldBind(&person) == nil {
		log.Println(person.Name)
		log.Println(person.Address)
		log.Println(person.Birthday)
	}
	// 以 String 回傳 Response。
	c.String(http.StatusOK, "Success\n")
	c.JSON(http.StatusOK, gin.H{"Name": person.Name, "Address": person.Address, "Birthday": person.Birthday})
}

// Get:
// $> curl -X GET "http://localhost:8085/testGet?name=Irara&address=Taipei&birthday=1987-08-23"
/*
{"Address":"Taipei","Birthday":"1987-08-23T00:00:00Z","Name":"Irara"}
*/

// Post:
// $> curl -X POST http://localhost:8085/testPost -F "name=Irara" -F "address=Taipei" -F "birthday=1987-08-23" -H "Content-Type:multipart/form-data"
/*
{"Address":"Taipei","Birthday":"1987-08-23T00:00:00Z","Name":"Irara"}
*/
