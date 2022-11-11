package request

/*
用户登录参数
*/

type LoginParams struct {
	Name     string `json:"name" binding:"required,checkname"`
	Password string `json:"password" binding:"required"`
}

/*
	用户注册参数
*/

type RegisterParams struct {
	Username string  `json:"username" binding:"required"`
	Name     string  `json:"name"`
	Password string  `json:"password" binding:"required"`
	Sex      *uint64 `json:"sex"`
	Phone    string  `json:"phone"`
}
