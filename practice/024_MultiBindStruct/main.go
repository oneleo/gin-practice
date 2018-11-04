package main

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type formA struct {
	Foo string `form:"foo" json:"foo" xml:"foo" binding:"required"`
}

type formB struct {
	Bar string `form:"bar" json:"bar" xml:"bar" binding:"required"`
}

func SomeHandler(c *gin.Context) {
	objA := formA{}
	objB := formB{}
	// This c.ShouldBind consumes c.Request.Body and it cannot be reused.
	if errA := c.ShouldBind(&objA); errA == nil {
		c.String(http.StatusOK, `the body should be formA`)
		// Always an error is occurred by this because c.Request.Body is EOF now.
	} else if errB := c.ShouldBind(&objB); errB == nil {
		c.String(http.StatusOK, `the body should be formB`)
	} else {
		c.String(http.StatusNoContent, "NoContent")
	}
}
func main() {
	r := gin.Default()
	r.GET("/", SomeHandler)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/?foo=hello"
/*
the body should be formA
*/

// $> curl -X GET "http://localhost:8085/?bar=world"
/*
the body should be formB
*/

// $> curl -X GET "http://localhost:8085/?foo=hello&bar=world"
/*
the body should be formA
*/
