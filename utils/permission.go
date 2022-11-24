package utils

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
)

func GetUserPermission(user model.User) ([]model.Permission, error) {
	var permissions []model.Permission

	err := global.DB.Preload("Roles.Permissions").First(&user).Error
	if err != nil {
		return permissions, err
	}

	for _, v := range user.Roles {
		permissions = append(permissions, v.Permissions...)
	}

	return permissions, nil
}
