package api

import (
	"gin_project/api/v1"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (l *LoginApi) LoginRouters(Router *gin.RouterGroup) {
	r := Router.Group("user") //.Use(middleware.TokenMiddleware())
	{
		r.POST("/login", api.ApiGroupApp.Login)           // 登录接口
		r.POST("/register", api.ApiGroupApp.Register)     //注册
		r.POST("/getCaptcha", api.ApiGroupApp.GetCaptcha) // 获取验证码
	}
}
