package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

// 模擬一些私有數據
var secrets = gin.H{
	"foo":    gin.H{"email": "foo@bar.com", "phone": "123433"},
	"austin": gin.H{"email": "austin@example.com", "phone": "666"},
	"lena":   gin.H{"email": "lena@guapa.com", "phone": "523443"},
}

// 使用認證中間件
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

	// Group using gin.BasicAuth() middleware
	// gin.Accounts is a shortcut for map[string]string
	// Group函數註冊了一個群組路由
	// gin.BasicAuth就是中間件，它的參數gin.Accounts其實是一個map[string]string類型的映射，這裡是用來記錄用戶名和密碼。
	authorized := r.Group("/admin", gin.BasicAuth(gin.Accounts{
		"foo":    "bar",
		"austin": "1234",
		"lena":   "hello2",
		"manu":   "4321",
	}))

	// /admin/secrets endpoint
	// hit "localhost:8080/admin/secrets
	// 然後在路由/admin之下，又註冊了/secrets路由，所以他的完整路由應該是/admin/secrets。處理器是getSecrets函數。
	authorized.GET("/secrets", func(c *gin.Context) {
		// get user, it was set by the BasicAuth middleware
		// 從上下文gin.Context中獲取用戶名。gin.AuthUserKey是一個字符串常量，定義如下：
		// const AuthUserKey = "user"
		// 這裡獲取的用戶名就是你訪問這個URL時輸入的用戶名，後台會驗證密碼，如果用戶名和密碼都對上了，認證才會成功。
		user := c.MustGet(gin.AuthUserKey).(string)

		// 接下來會利用獲得的用戶名去secrets結構中查找用戶信息。gin.H是一個map[string]interface{}的映射。最後通過JSON返回查詢結果。
		if secret, ok := secrets[user]; ok {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": secret})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "secret": "NO SECRET :("})
		}
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// 使用瀏覽器瀏覽：
// http://localhost:8085/admin/secrets

// $> curl --head -X GET "http://localhost:8085/admin/secrets"
/*
HTTP/1.1 401 Unauthorized
Www-Authenticate: Basic realm="Authorization Required"
Date: Sat, 03 Nov 2018 12:18:02 GMT
Content-Length: 0
*/

// $> curl -X GET "http://foo:bar@localhost:8085/admin/secrets"
/*
{"secret":{"email":"foo@bar.com","phone":"123433"},"user":"foo"}
*/

// $> curl -X GET "http://austin:1234@localhost:8085/admin/secrets"
/*
{"secret":{"email":"austin@example.com","phone":"666"},"user":"austin"}
*/

// $> curl -X GET "http://lena:hello2@localhost:8085/admin/secrets"
/*
{"secret":{"email":"lena@guapa.com","phone":"523443"},"user":"lena"}
*/

// $> curl -X GET "http://manu:4321@localhost:8085/admin/secrets"
/*
{"secret":"NO SECRET :(","user":"manu"}
*/
