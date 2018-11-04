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

	// Default With the Logger and Recovery middleware already attached
	// 默認情況已啟用了log和恢復中間件
	//r := gin.Default()

	// 默認情況下創建一個沒有任何中間件的路由器
	r := gin.New()

	// Global middleware
	// Logger middleware will write the logs to gin.DefaultWriter even if you set with GIN_MODE=release.
	// By default gin.DefaultWriter = os.Stdout
	r.Use(gin.Logger())

	// Recovery middleware recovers from any panics and writes a 500 if there was one.
	r.Use(gin.Recovery())

	// 每個路由中,你可以使用任意多個中間件.
	r.GET("/benchmark", MyBenchLogger, benchEndpoint)

	// 權限組
	// authorized := r.Group("/", AuthRequired())
	// 等同於:
	authorized := r.Group("/")
	// 在這組路由中,我們使用自定義的中間件
	// AuthRequired() 中間件只在 "authorized" 組中使用.
	authorized.Use(AuthRequired)
	{
		authorized.POST("/login", loginEndpoint)
		authorized.POST("/submit", submitEndpoint)
		authorized.POST("/read", readEndpoint)

		// 嵌套組
		testing := authorized.Group("/testing")
		testing.GET("/analytics", analyticsEndpoint)
	}

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

func MyBenchLogger(c *gin.Context) {
	log.Println("exec MyBenchLogger")
	c.String(http.StatusOK, "exec MyBenchLogger\n")
	//你可以寫一些邏輯代碼

	// 執行該中間件之後的邏輯
	c.Next()
}

func AuthRequired(c *gin.Context) {
	log.Println("exec AuthRequired")
	c.String(http.StatusOK, "exec AuthRequired\n")
	//你可以寫一些邏輯代碼

	// 執行該中間件之後的邏輯
	c.Next()
}

func benchEndpoint(c *gin.Context) {
	log.Printf("URL: %s\n", c.Request.RequestURI)
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}
func loginEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}
func submitEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}
func readEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}

func analyticsEndpoint(c *gin.Context) {
	c.String(http.StatusOK, "URL: %s\n", c.Request.RequestURI)
}

// $> curl -X GET "http://localhost:8085/benchmark"
/*
exec MyBenchLogger
Hello
*/

// $> curl -X POST "http://localhost:8085/login"
/*
exec AuthRequired
URL: /login
*/

// $> curl -X POST "http://localhost:8085/submit"
/*
exec AuthRequired
URL: /submit
*/

// $> curl -X POST "http://localhost:8085/read"
/*
exec AuthRequired
URL: /read
*/

// $> curl -X GET "http://localhost:8085/testing/analytics"
/*
exec AuthRequired
URL: /testing/analytics
*/
