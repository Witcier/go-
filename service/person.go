package service

import (
	"errors"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils"
	"witcier/go-api/utils/category"

	"github.com/gin-gonic/gin"
)

type PersonService struct{}

func (s *PersonService) PersonDetail(c *gin.Context) (model.User, error) {
	user, err := utils.GetAuthUser(c)
	if err != nil {
		return user, err
	}

	return user, nil
}

func (s *PersonService) PersonUpdatePassword(c *gin.Context, req request.PersonUpdatePassword) error {
	user, err := utils.GetAuthUser(c)
	if err != nil {
		return err
	}

	if ok := utils.BcryptCheck(req.OldPassword, user.Password); !ok {
		return errors.New("密码错误")
	}

	return global.DB.Updates(&model.User{
		ModelID: global.ModelID{
			ID: user.ID,
		},
		Password: utils.BcryptHash(req.Password),
	}).Error
}

func (s *PersonService) Menu(c *gin.Context) (response.CategoryResult, error) {
	var resp response.CategoryResult
	var menus []model.Menu

	user, err := utils.GetAuthUser(c)
	if err != nil {
		return resp, err
	}

	err = global.DB.Preload("Roles.Menus").First(&user).Error
	if err != nil {
		return resp, err
	}

	for _, v := range user.Roles {
		menus = append(menus, v.Menus...)
	}

	resp.List = category.MenuCategory(menus, 0)

	return resp, nil
}
