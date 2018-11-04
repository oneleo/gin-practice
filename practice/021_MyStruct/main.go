package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

type StructA struct {
	FieldA string `form:"field_a"`
}

type StructB struct {
	NestedStruct StructA
	FieldB       string `form:"field_b"`
}

type StructC struct {
	NestedStructPointer *StructA
	FieldC              string `form:"field_c"`
}

type StructD struct {
	NestedAnonyStruct struct {
		FieldX string `form:"field_x"`
	}
	FieldD string `form:"field_d"`
}

// NOTE: NOT support the follow style struct:
type StructX struct {
	X struct{} `form:"name_x"` // HERE have form
}

// NOTE: NOT support the follow style struct:
type StructY struct {
	Y StructX `form:"name_y"` // HERE hava form
}

// NOTE: NOT support the follow style struct:
type StructZ struct {
	Z *StructZ `form:"name_z"` // HERE hava form
}

func GetDataB(c *gin.Context) {
	var b StructB
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStruct,
		"b": b.FieldB,
	})
}

func GetDataC(c *gin.Context) {
	var b StructC
	c.Bind(&b)
	c.JSON(200, gin.H{
		"a": b.NestedStructPointer,
		"c": b.FieldC,
	})
}

func GetDataD(c *gin.Context) {
	var b StructD
	c.Bind(&b)
	c.JSON(200, gin.H{
		"x": b.NestedAnonyStruct,
		"d": b.FieldD,
	})
}

func main() {
	r := gin.Default()
	r.GET("/getb", GetDataB)
	r.GET("/getc", GetDataC)
	r.GET("/getd", GetDataD)

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// $> curl -X GET "http://localhost:8085/getb?field_a=hello&field_b=world"
/*
{"a":{"FieldA":"hello"},"b":"world"}
*/

// $> curl -X GET "http://localhost:8085/getc?field_a=hello&field_c=world"
/*
{"a":{"FieldA":"hello"},"c":"world"}
*/

// $> curl -X GET "http://localhost:8085/getd?field_x=hello&field_d=world"
/*
{"d":"world","x":{"FieldX":"hello"}}
*/
