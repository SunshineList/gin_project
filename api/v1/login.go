package api

import (
	"github.com/gin-gonic/gin"
	"go_project/gin_project/gin_test/model"
	"go_project/gin_project/gin_test/model/request"
	"go_project/gin_project/gin_test/response"
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
	} else {
		response.OkAndData(loginParams, context)
	}
}

func (l *LoginApi) Register(context *gin.Context) {
	var registerParams = request.RegisterParams{}
	err := context.ShouldBindJSON(&registerParams)
	if err != nil {
		response.FailAndMsg(err.Error(), context)
	}
	user := model.User{
		Username: registerParams.Username,
		Name:     registerParams.Name,
		Password: registerParams.Password,
		Sex:      registerParams.Sex,
	}

	r, err := userService.RegisterService(user)

	if err != nil {
		response.FailAndMsg("注册失败", context)
	}
	response.OkAndData(r, context)
}
