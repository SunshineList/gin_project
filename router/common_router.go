package api

import (
	"github.com/gin-gonic/gin"
	"go_project/gin_project/gin_test/api/v1"
)

type UploadApi struct{}

func (u *UploadApi) UploadRouters(Router *gin.RouterGroup) {
	r := Router.Group("common")
	{
		r.POST("/upload_file", api.ApiGroupApp.UploadFile)
	}
}
