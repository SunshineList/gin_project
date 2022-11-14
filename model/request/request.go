package request

/*
用户登录参数
*/

type LoginParams struct {
	Username   string `json:"username" binding:"required"`
	Password   string `json:"password" binding:"required"`
	CaptchaId  string `json:"captcha_id" binding:"required"`
	CaptchaVal string `json:"captcha_val" binding:"required"`
}

/*
	用户注册参数
*/

type RegisterParams struct {
	Username  string  `json:"username" binding:"required"`
	Name      string  `json:"name"`
	Password  string  `json:"password" binding:"required"`
	Rpassword string  `json:"rpassword" binding:"required"`
	Sex       *uint64 `json:"sex"`
	Phone     string  `json:"phone"`
}
