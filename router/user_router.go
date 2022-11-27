package api

import (
	"gin_project/api/v1"
	"gin_project/middleware"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (l *LoginApi) LoginRouters(Router *gin.RouterGroup) {
	r := Router.Group("user") //.Use(middleware.TokenMiddleware())
	{
		r.POST("/login", api.ApiGroupApp.Login)                                            // 登录接口
		r.POST("/register", api.ApiGroupApp.Register)                                      //注册
		r.GET("/getCaptcha", api.ApiGroupApp.GetCaptcha)                                   // 获取验证码
		r.GET("/getUserInfo", middleware.JWTAuthMiddleware(), api.ApiGroupApp.GetUserInfo) // 查看当前用户信息
		r.GET("/index", api.ApiGroupApp.TestHtml)
	}
}
