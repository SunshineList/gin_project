package api

import (
	"gin_project/api/v1"
	"github.com/gin-gonic/gin"
)

type UploadApi struct{}

func (u *UploadApi) UploadRouters(Router *gin.RouterGroup) {
	r := Router.Group("common")
	{
		r.POST("/upload_file", api.ApiGroupApp.UploadFile)
	}
}
