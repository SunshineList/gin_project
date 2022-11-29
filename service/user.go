package service

import (
	"errors"
	"gin_project/common/config"
	"gin_project/model"
	"gin_project/model/request"
	"gin_project/response"
	"gin_project/utils"
	"gorm.io/gorm"
)

type UserService struct{}

// 注册

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

//登录

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

// 获取当前登录用户信息

func (userService *UserService) GetUserById(Id any) (userInt interface{}, err error) {
	var user model.User
	if errors.Is(config.GVA_DB.Where("id = ?", Id).First(&user).Error, gorm.ErrRecordNotFound) {
		return nil, errors.New("未找到用户")
	}
	res, err := response.ToResponse(user, request.UserInfo{})
	return res, err
}

// 修改用户信息

func (userService *UserService) UpdateUserService(u model.User) (data interface{}, err error) {
	var user model.User
	err = config.GVA_DB.Updates(&u).Error
	config.GVA_DB.Where("id = ?", u.ID).First(&user)
	res, err := response.ToResponse(user, request.UserInfo{})
	return res, err
}
