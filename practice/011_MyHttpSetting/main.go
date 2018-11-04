package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

// 自定義HTTP配置
// Use http.ListenAndServe() directly, like this:
func main() {
	router := gin.Default()
	http.ListenAndServe(":8085", router)
}
