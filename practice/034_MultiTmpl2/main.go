package main

import (
	"path/filepath"

	"github.com/gin-contrib/multitemplate"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.HTMLRender = loadTemplates("./templates")
	router.GET("/", func(c *gin.Context) {
		c.HTML(200, "index.html", gin.H{
			"title": "Welcome!",
		})
	})
	router.GET("/article", func(c *gin.Context) {
		c.HTML(200, "article.html", gin.H{
			"title": "Html5 Article Engine",
		})
	})

	router.Run(":8085")
}

func loadTemplates(templatesDir string) multitemplate.Renderer {
	r := multitemplate.NewRenderer()

	layouts, err := filepath.Glob(templatesDir + "/layouts/*.html")
	if err != nil {
		panic(err.Error())
	}

	includes, err := filepath.Glob(templatesDir + "/includes/*.html")
	if err != nil {
		panic(err.Error())
	}

	// Generate our templates map from our layouts/ and includes/ directories
	for _, include := range includes {
		layoutCopy := make([]string, len(layouts))
		copy(layoutCopy, layouts)
		files := append(layoutCopy, include)
		r.AddFromFiles(filepath.Base(include), files...)
	}
	return r
}

// 使用瀏覽器瀏覽：
// http://localhost:8085/
// http://localhost:8085/article
