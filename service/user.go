package service

import (
	"errors"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils"

	"gorm.io/gorm"
)

type UserService struct{}

func (userService *UserService) ListUser(r request.List) (response.PageResult, error) {
	var resp response.PageResult
	var total int64

	limit := r.PageSize
	offset := r.PageSize * (r.Page - 1)
	db := global.DB.Model(&model.User{})
	var users []model.User
	err := db.Count(&total).Error
	if err != nil {
		return resp, err
	}

	err = db.Limit(limit).Offset(offset).Find(&users).Error

	resp = response.PageResult{
		List:     users,
		Total:    total,
		Page:     r.Page,
		PageSize: r.PageSize,
	}

	return resp, err
}

func (userService *UserService) StoreUser(r request.StoreUser) (err error) {
	var user model.User
	if !errors.Is(global.DB.Where("username = ?", r.Username).First(&user).Error, gorm.ErrRecordNotFound) {
		return errors.New("用户已经注册")
	}

	newUser := &model.User{
		Username:     r.Username,
		RealName:     r.RealName,
		Email:        r.Email,
		Phone:        r.Phone,
		Sex:          r.Sex,
		Birth:        r.Birth,
		EntryTime:    r.EntryTime,
		Status:       r.Status,
		IsAdmin:      r.IsAdmin,
		DepartmentID: r.DepartmentID,
		PositionID:   r.PositionID,
	}

	newUser.Password = utils.BcryptHash(global.Config.System.UserDefaultPassword)

	err = global.DB.Create(&newUser).Error

	return err
}

func (userService *UserService) UpdateUser(id uint, r request.UpdateUser) error {

	user := &model.User{
		ModelID: global.ModelID{
			ID: id,
		},
		Username:     r.Username,
		RealName:     r.RealName,
		Email:        r.Email,
		Phone:        r.Phone,
		Sex:          r.Sex,
		Birth:        r.Birth,
		EntryTime:    r.EntryTime,
		Status:       r.Status,
		IsAdmin:      r.IsAdmin,
		DepartmentID: r.DepartmentID,
		PositionID:   r.PositionID,
	}

	return global.DB.Updates(&user).Error
}

func (userService *UserService) DeleteUser(id uint) error {
	var user model.User
	err := global.DB.Where("id = ?", id).Delete(&user).Error
	if err != nil {
		return err
	}

	return nil
}
