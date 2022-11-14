package api

import (
	"gin_project/model"
	"gin_project/model/request"
	"gin_project/response"
	"gin_project/utils"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

func (l *LoginApi) Login(context *gin.Context) {

	var loginParams = request.LoginParams{}
	err := context.ShouldBindJSON(&loginParams)
	if err != nil {
		response.FailAndMsg(err.Error(), context)
		return
	}
	user := &model.User{
		Username: loginParams.Username,
		Password: loginParams.Password,
	}

	if !utils.CheckCaptcha(loginParams.CaptchaId, loginParams.CaptchaVal) {
		response.FailAndMsg("验证码错误", context)
		return
	}

	res, err := userService.LoginService(*user)

	if err != nil {
		response.FailAndMsg(err.Error(), context)
		return
	}
	response.OkAndData(res, "登录成功", context)
}

func (l *LoginApi) Register(context *gin.Context) {
	var registerParams = request.RegisterParams{}
	err := context.ShouldBindJSON(&registerParams)
	if err != nil {
		response.FailAndMsg(utils.Translate(err), context)
		return
	}

	if registerParams.Password != registerParams.Rpassword {
		response.FailAndMsg("两次密码输入不一致", context)
		return
	}

	user := &model.User{
		Username: registerParams.Username,
		Name:     registerParams.Name,
		Password: registerParams.Password,
		Sex:      registerParams.Sex,
		Phone:    registerParams.Phone,
	}

	r, err := userService.RegisterService(*user)

	if err != nil {
		response.FailAndMsg("注册失败", context)
		return
	}
	response.OkAndData(r, "注册成功", context)
}
