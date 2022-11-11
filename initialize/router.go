package initialize

import (
	"github.com/gin-gonic/gin"
	router "go_project/gin_project/gin_test/router"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 20 << 20 // 设置最大上传20M
	LoginPath := r.Group("/v1")
	{
		router.RouterGroupApp.LoginRouters(LoginPath)
		router.RouterGroupApp.UploadRouters(LoginPath)
	}
	return r
}
