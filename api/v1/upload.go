package api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type UploadApi struct{}

func (u *UploadApi) UploadFile(context *gin.Context) {
	form, _ := context.MultipartForm()
	files := form.File["upload[]"]

	fmt.Println(files)

	for _, file := range files {
		log.Println(file.Filename)

		// 上传文件到指定的路径
		// c.SaveUploadedFile(file, dst)
	}
}
