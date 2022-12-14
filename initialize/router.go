package initialize

import (
	router "gin_project/router"
	"github.com/gin-gonic/gin"
	"net/http"
)

func InitRoutes() *gin.Engine {
	r := gin.Default()
	r.MaxMultipartMemory = 20 << 20 // 设置最大上传20M
	r.StaticFS("/media", http.Dir("./media"))
	r.LoadHTMLGlob("templates/*")
	LoginPath := r.Group("/v1")
	{
		router.RouterGroupApp.LoginRouters(LoginPath)
		router.RouterGroupApp.UploadRouters(LoginPath)
	}
	return r
}
