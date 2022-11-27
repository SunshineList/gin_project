package api

import (
	"gin_project/common/config"
	"gin_project/middleware"
	"gin_project/model"
	"gin_project/model/request"
	"gin_project/response"
	"gin_project/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt"
	"net/http"
)

type LoginApi struct{}

func (l *LoginApi) Login(context *gin.Context) {

	var loginParams = request.LoginParams{}
	err := context.ShouldBindJSON(&loginParams)
	if err != nil {
		response.FailAndMsg(utils.Translate(err), context)
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

	if res.Status != model.ACTIVE {
		response.FailAndMsg("账号已冻结请联系管理员", context)
		return
	}

	// 签发jwt
	jwtToken, _ := utils.JwtToken(utils.MyCustomClaims{Id: res.ID, Username: res.Username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: config.JWT_EXPITE_TIME, // 过期时间
		}})

	response.OkAndData(map[string]interface{}{
		"username": res.Username,
		"name":     res.Name,
		"token":    jwtToken,
	}, "登录成功", context)
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

func (l *LoginApi) GetUserInfo(context *gin.Context) {
	user, err := middleware.CurrentUser(context)
	if err != nil {
		response.FailAndMsg(err.Error(), context)
		return
	}
	println(user)
	response.OkAndData(user, "查询成功", context)
}

func (l *LoginApi) TestHtml(context *gin.Context) {
	context.HTML(http.StatusOK, "index.html", gin.H{
		"name": "hhhhhhhhh",
	})
}
