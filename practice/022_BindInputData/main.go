package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 定義的 Login 結構體
// 該 struct 可以綁定在 Form 和 JSON 中
// binding:"required" 意思是必要參數。如果未提供，Bind 會返回 error
type Login struct {
	User     string `form:"user" json:"user" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
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

	// Example for binding JSON ({"user": "manu", "password": "123"})
	// POST 到這個路由一段 JSON, 如 ({"user": "manu", "password": "123"})
	r.POST("/loginJSON", func(c *gin.Context) {
		var json Login
		// 驗證數據並綁定
		if err := c.ShouldBindJSON(&json); err == nil {
			if json.User == "manu" && json.Password == "123" {
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// Example for binding a HTML form (user=manu&password=123)
	// POST 到這個路由一個 Form 表單 (user=manu&password=123)
	r.POST("/loginForm", func(c *gin.Context) {
		var form Login
		// 驗證數據並綁定
		if err := c.ShouldBind(&form); err == nil {
			if form.User == "manu" && form.Password == "123" {
				// 以 Json 回傳 Response。
				c.JSON(http.StatusOK, gin.H{"status": "you are logged in"})
			} else {
				// 以 Json 回傳 Response。
				c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
			}
		} else {
			// 以 Json 回傳 Response。
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		}
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// Json Input:
// $> curl --verbose -X POST "http://localhost:8085/loginJSON" -d "{ \"user\": \"manu\", \"password\": \"123\" }" -H "Content-Type:application/json"
/*
{"status":"you are logged in"}
*/

// $> curl --verbose -X POST http://localhost:8085/loginJSON -d "{ \"user\": \"manu\", \"password\": \"567\" }" -H "Content-Type:application/json"
/*
{"status":"unauthorized"}
*/

// Form Input:
// $> curl --verbose -X POST "http://localhost:8085/loginForm" -F "user=manu" -F "password=123"
/*
{"status":"you are logged in"}
*/

// $> curl --verbose -X POST "http://localhost:8085/loginForm" -F "user=manu" -F "password=456"
/*
{"status":"unauthorized"}
*/
