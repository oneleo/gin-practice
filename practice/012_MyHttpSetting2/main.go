package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// 自定義HTTP配置
// Use http.ListenAndServe() directly, like this:
func main() {
	router := gin.Default()

	s := &http.Server{
		Addr:           ":8085",
		Handler:        router,
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}
	s.ListenAndServe()
}
