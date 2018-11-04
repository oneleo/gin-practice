package main

import (
	"html/template"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.New()
	t, err := loadTemplate()
	if err != nil {
		panic(err)
	}
	r.SetHTMLTemplate(t)
	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/index.tmpl", gin.H{
			"Foo": "World",
		})
	})
	r.GET("/bar", func(c *gin.Context) {
		c.HTML(http.StatusOK, "/html/bar.tmpl", gin.H{
			"Bar": "World",
		})
	})
	r.Run(":8085")
}

func loadTemplate() (*template.Template, error) {
	t := template.New("")
	for name, file := range Assets.Files {
		if file.IsDir() || !strings.HasSuffix(name, ".tmpl") {
			continue
		}
		h, err := ioutil.ReadAll(file)
		if err != nil {
			return nil, err
		}
		t, err = t.New(name).Parse(string(h))
		if err != nil {
			return nil, err
		}
	}
	return t, nil
}

// 將 main.go 包成 bin 可執行檔時，因為執行檔並不包含 html/tmpl 內的檔案
// 所以需要 jessevdk/go-assets-builder 專案，將 html/tmpl 包成 assets.go 檔
// 如此就可以成為單一執行檔。
// $> go get -u -v github.com/jessevdk/go-assets
// $> go get -u -v github.com/jessevdk/go-assets-builder
//
// $> cd ./046_BinaryTmpl/原始檔
// $> go-assets-builder ./html -o ./assets.go
// $> go build -o assets-in-binary
