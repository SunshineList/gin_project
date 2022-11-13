package service

import (
	"errors"
	"gin_project/common/config"
	"gin_project/model"
	"gin_project/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func (u *UserService) RegisterService(user model.User) (userInt model.User, err error) {

	// 使用errors.Is 判断前后error是否一致
	if !errors.Is(config.GVA_DB.Where("username = ?", user.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return user, errors.New("用户已注册")
	}
	// 处理密码
	user.Password = utils.MD5V([]byte(user.Password))
	err = config.GVA_DB.Create(&user).Error

	return user, err
}
