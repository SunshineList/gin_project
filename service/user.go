package service

import (
	"errors"
	"gin_project/common/config"
	"gin_project/model"
	"gin_project/utils"
	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) RegisterService(user model.User) (userInt *model.User, err error) {

	// 使用errors.Is 判断前后error是否一致
	if !errors.Is(config.GVA_DB.Where("username = ?", user.Username).First(&model.User{}).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户已注册")
	}

	// 处理密码
	user.Password = utils.MD5V([]byte(user.Password))
	err = config.GVA_DB.Create(&user).Error

	return &user, err
}

func (userService *UserService) LoginService(u model.User) (userInt *model.User, err error) {
	var user model.User
	if errors.Is(config.GVA_DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("用户未注册请先注册")
	}
	if !utils.Md5Check([]byte(u.Password), user.Password) {
		return nil, errors.New("密码错误")
	}
	return &user, nil
}

func (userService *UserService) GetUserById(Id any) (userInt *model.User, err error) {
	var user model.User
	if errors.Is(config.GVA_DB.Where("id = ?", Id).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("未找到用户")
	}
	return &user, err
}
