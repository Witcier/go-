package initialize

import (
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"go.uber.org/zap"
)

func GenerateTokenForTest() {
	var user model.User

	j := utils.NewJWT()
	global.DB.First(&user)

	claims := j.CreateClaims(request.BaseClaims{
		ID:       user.ID,
		Username: user.Username,
		RealName: user.RealName,
	})

	token, err := j.CreateToken(claims)
	if err != nil {
		global.Log.Error("create token fail", zap.Error(err))
	}

	global.Log.Info("token for test:" + token)
}
