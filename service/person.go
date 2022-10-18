package service

import (
	"errors"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

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
