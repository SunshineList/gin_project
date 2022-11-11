package api

import (
	"github.com/gin-gonic/gin"
	"go_project/gin_project/gin_test/api/v1"
	"go_project/gin_project/gin_test/middleware"
)

type LoginApi struct{}

func (l *LoginApi) LoginRouters(Router *gin.RouterGroup) {
	r := Router.Group("user").Use(middleware.TokenMiddleware())
	{
		r.POST("/login", api.ApiGroupApp.Login)       // 登录接口
		r.POST("/register", api.ApiGroupApp.Register) //注册
	}
}
