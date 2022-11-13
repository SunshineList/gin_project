package api

import (
	"gin_project/model"
	"gin_project/model/request"
	"gin_project/response"
	"github.com/gin-gonic/gin"
)

type LoginApi struct{}

//func checkName(fl validator.FieldLevel) bool {
//	value := fl.Field().Interface().(string) // 反射拿到返回值
//	if value != "admin" {
//		return false
//	}
//	return true
//}

func (l *LoginApi) Login(context *gin.Context) {
	//if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
	//	v.RegisterValidation("checkname", checkName)
	//}

	var loginParams = request.LoginParams{}
	err := context.ShouldBindJSON(&loginParams)
	if err != nil {
		response.FailAndMsg(err.Error(), context)
		return
	}
	response.OkAndData(loginParams, context)
}

func (l *LoginApi) Register(context *gin.Context) {
	var registerParams = request.RegisterParams{}
	err := context.ShouldBindJSON(&registerParams)
	if err != nil {
		response.FailAndMsg(err.Error(), context)
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
	response.OkAndData(r, context)
}
