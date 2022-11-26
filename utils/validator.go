package utils

import (
	"github.com/mojocn/base64Captcha"
)

/**
验证器相关代码
*/

// 验证码校验

func CheckCaptcha(captchaId string, captchaVal string) bool {
	var store = base64Captcha.DefaultMemStore
	return store.Verify(captchaId, captchaVal, true)
}

// 性别转换

func ChangeSex(sex string) string {
	if sex == "1" {
		return "男"
	}
	return "女"
}

// 状态转换

func ChangeStatus(status string) string {
	if status == "1" {
		return "正常"
	}
	return "异常"
}
