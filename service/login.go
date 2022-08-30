package service

import (
	"errors"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils"

	"go.uber.org/zap"
)

type LoginService struct{}

func (s *LoginService) Login(req request.Login) (response.LoginResponse, error) {
	var user model.User
	var resp response.LoginResponse

	err := global.DB.Where("username = ?", req.Username).First(&user).Error

	if err == nil {
		if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
			return resp, errors.New("密码错误")
		}
	}

	j := &utils.JWT{Signingkey: []byte(global.Config.JWT.SigningKey)}
	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		Username: user.Username,
		RealName: user.RealName,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.Log.Error("create token fail", zap.Error(err))

		return resp, err
	}

	resp = response.LoginResponse{
		User:      user,
		Type:      global.Config.JWT.TokenType,
		ExpiresAt: claims.ExpiresAt.Unix(),
		Token:     token,
	}

	return resp, err
}
