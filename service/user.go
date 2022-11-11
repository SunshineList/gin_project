package service

import (
	"errors"
	"go_project/gin_project/gin_test/common"
	"go_project/gin_project/gin_test/model"
)

type UserService struct{}

func (u *UserService) RegisterService(user model.User) (userInt model.User, err error) {

	if e := common.GVA_DB.Where("username = ?", user.Username).First(&model.User{}).Error; e != nil {
		return user, errors.New("用户已注册")
	}
	err = common.GVA_DB.Create(&user).Error
	// 处理密码
	return user, err
}
