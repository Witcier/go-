package service

import (
	"errors"
	"time"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/model/response"
	"witcier/go-api/utils"

	"github.com/golang-jwt/jwt/v4"
	"go.uber.org/zap"
)

type LoginService struct{}

var j = utils.NewJWT()

func (s *LoginService) Login(req request.Login) (response.LoginResponse, error) {
	var user model.User
	var resp response.LoginResponse

	err := global.DB.Where("username = ?", req.Username).First(&user).Error
	if err != nil {
		return resp, err
	}

	if ok := utils.BcryptCheck(req.Password, user.Password); !ok {
		return resp, errors.New("密码错误")
	}

	// j := &utils.JWT{Signingkey: []byte(global.Config.JWT.SigningKey)}
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

func (s *LoginService) Refresh(oldToken string) (response.LoginResponse, error) {
	var user model.User
	var resp response.LoginResponse

	claims, err := j.ParseToken(oldToken)
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Hour * global.Config.JWT.ExpiresHour))

	err = global.DB.Where("ID = ?", claims.BaseClaims.ID).First(&user).Error
	if err != nil {
		return resp, err
	}

	newToken, err := j.RefreshToken(oldToken, *claims)
	if err != nil {
		return resp, err
	}

	resp = response.LoginResponse{
		User:      user,
		Type:      global.Config.JWT.TokenType,
		ExpiresAt: claims.ExpiresAt.Unix(),
		Token:     newToken,
	}

	return resp, nil
}
