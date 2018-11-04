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

	// gin.H 本質是 map[string]interface{}
	r.GET("/someJSON", func(c *gin.Context) {
		// 會輸出頭格式為 application/json; charset=UTF-8 的 json 字符串
		c.JSON(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/moreJSON", func(c *gin.Context) {
		// 直接使用結構體定義
		var msg struct {
			Name    string `json:"user"`
			Message string
			Number  int
		}
		msg.Name = "Lena"
		msg.Message = "hey"
		msg.Number = 123
		// 會輸出  {"user": "Lena", "Message": "hey", "Number": 123}
		c.JSON(http.StatusOK, msg)
	})

	r.GET("/someXML", func(c *gin.Context) {
		// 會輸出頭格式為 text/xml; charset=UTF-8 的 xml 字符串
		c.XML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	r.GET("/someYAML", func(c *gin.Context) {
		// 會輸出頭格式為 text/yaml; charset=UTF-8 的 yaml 字符串
		c.YAML(http.StatusOK, gin.H{"message": "hey", "status": http.StatusOK})
	})

	// 使用SecureJSON來防止json劫持. 如果返回的結果是數組則會在返回數據前添加默認的前綴 "while(1),".
	// 你也可以自定義你自己的安全json前綴
	// r.SecureJsonPrefix(")]}',\n")
	r.GET("/secureJSON", func(c *gin.Context) {
		names := []string{"lena", "austin", "foo"}

		// Will output  :   while(1);["lena","austin","foo"]
		c.SecureJSON(http.StatusOK, names)
	})

	// 使用 JSONP 從不同域中的服務器請求數據。如果查詢參數回調存在，請將回調添加到響應主體。
	// 這裡假定訪問的url地址是 /JSONP?callback=x
	// url中必須存在callback= 才會返回jsonp,否則返回json
	r.GET("/JSONP", func(c *gin.Context) {
		data := map[string]interface{}{
			"foo": "bar",
		}

		//callback is x
		// Will output  :   x({\"foo\":\"bar\"})
		c.JSONP(http.StatusOK, data)
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/someJSON"
/*
{"message":"hey","status":200}
*/

// $> curl -X GET "http://localhost:8085/moreJSON"
/*
{"user":"Lena","Message":"hey","Number":123}
*/

// $> curl -X GET "http://localhost:8085/someXML"
/*
<map><message>hey</message><status>200</status></map>
*/

// $> curl -X GET "http://localhost:8085/someYAML"
/*
message: hey
status: 200
*/

// $> curl -X GET "http://localhost:8085/someYAML"
/*
message: hey
status: 200
*/

// $> curl -X GET "http://localhost:8085/secureJSON"
/*
while(1);["lena","austin","foo"]
*/

// $> curl -X GET "http://localhost:8085/JSONP?callback=xyz"
/*
xyz({"foo":"bar"})
*/

// $> curl -X GET "http://localhost:8085/JSONP?call=xyz"
/*
{"foo":"bar"}
*/
