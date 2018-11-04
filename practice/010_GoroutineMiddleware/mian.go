package main

import (
	"log"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
)

// 在中間件或處理程序中啟動新的Goroutines時, 你 一定不要 使用它內部的原始上下文, 你必須使用只讀副本.
func main() {
	r := gin.Default()
	// 異步執行
	r.GET("/long_async", func(c *gin.Context) {
		var wg sync.WaitGroup

		// create copy to be used inside the goroutine
		cCp := c.Copy()
		wg.Add(1)
		go func() {
			defer wg.Done()
			// 模擬一個耗時任務
			time.Sleep(time.Duration(5) * time.Second)

			// 一定要使用複製的cCp
			log.Println("Done! in path " + cCp.Request.URL.Path)
			cCp.String(http.StatusOK, "Done! in path %s\n", cCp.Request.URL.Path)
		}()
		log.Println("Still Waiting!")
		c.String(http.StatusOK, "Still Waiting!")
		wg.Wait()
	})

	// 同步執行
	r.GET("/long_sync", func(c *gin.Context) {
		// simulate a long task with time.Sleep(). 5 seconds
		time.Sleep(time.Duration(5) * time.Second)

		// 不使用Goroutines則不需要複製
		log.Println("Done! in path " + c.Request.URL.Path)
		c.String(http.StatusOK, "Done! in path %s\n", c.Request.URL.Path)
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/long_async"
/*
{"secret":{"email":"foo@bar.com","phone":"123433"},"user":"foo"}
*/

// $> curl -X GET "http://localhost:8085/long_sync"
/*
{"secret":{"email":"foo@bar.com","phone":"123433"},"user":"foo"}
*/
