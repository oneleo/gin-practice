package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.POST("/post", func(c *gin.Context) {
		// 獲取原始字節
		d, err := c.GetRawData()
		if err != nil {
			log.Fatalln(err)
		}
		log.Println(string(d))
		// 以 String 輸出 Response。
		c.String(200, "ok")
	})
	router.Run(":8080")
}

// $ curl -v -X POST \
//   http://localhost:8080/post \
//   -H 'content-type: application/json' \
//   -d '{ "user": "manu" }'

// > curl -X POST http://localhost:8080/post -d "{ \"user\": \"manu\" }" -H "Content-Type:multipart/form-data"
