package main

import (
	"fmt"
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

	// 初始化引擎。
	// 用默認的中間件創建一個gin路由器:
	// logger and recovery (crash-free) middleware
	// 記錄 恢復(不崩潰) 中間件
	r := gin.Default()

	// 設置文件上傳大小 router.MaxMultipartMemory = 8 << 20  // 8 MiB
	// 處理單一的文件上傳
	r.POST("/upload", func(c *gin.Context) {
		// 拿到這個文件
		file, _ := c.FormFile("file")
		log.Println(file.Filename)

		savePath := "./tmp/"
		dst := savePath + file.Filename

		// Upload the file to specific dst.
		err := c.SaveUploadedFile(file, dst)
		if err != nil {
			panic(err)
		}

		c.String(http.StatusOK, fmt.Sprintf("'%s' uploaded!", file.Filename))
	})

	// 處理多個文件的上傳
	r.POST("/uploads", func(c *gin.Context) {
		form, _ := c.MultipartForm()
		// 拿到集合
		files := form.File["upload[]"]
		savePath := "./tmp/"

		for _, file := range files {
			log.Println(file.Filename)
			dst := savePath + file.Filename
			// Upload the file to specific dst.
			c.SaveUploadedFile(file, dst)
		}
		// 會回傳到至目前頁面的訊息，以 String 輸出 Response。
		c.String(http.StatusOK, fmt.Sprintf("%d files uploaded!", len(files)))
	})

	// 指定端口號為 :8085
	err := r.Run(":8085")
	if err != nil {
		log.Fatalln(err)
	}
}

// 單一文件上傳
// $> curl -X POST "http://localhost:8085/upload" -F "file=@./README.md" -H "Content-Type:multipart/form-data"

// 多文件上傳
// $> curl -X POST "http://localhost:8085/uploads" -F "upload[]=@./document/010_UploadFile/test1.txt" -F "upload[]=@./document/010_UploadFile/test2.txt" -H "Content-Type:multipart/form-data"
