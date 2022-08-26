package service

import (
	"errors"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/utils"

	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) StoreUser(u model.User) (userInter model.User, err error) {
	var user model.User
	if !errors.Is(global.DB.Where("username = ?", u.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return userInter, errors.New("用户已经注册")
	}

	u.Password = utils.BcryptHash(global.Config.System.UserDefaultPassword)

	err = global.DB.Create(&u).Error

	return u, err
}

func (userService *UserService) UpdateUser(u model.User) error {
	return global.DB.Updates(&u).Error
}
