package v1

import (
	"witcier/go-api/global"
	"witcier/go-api/model/common/response"
	"witcier/go-api/model/request"
	resp "witcier/go-api/model/response"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
	"github.com/mojocn/base64Captcha"
	"go.uber.org/zap"
)

type LoginApi struct{}

var store = utils.NewCaptchaRedisStore()

func (a *LoginApi) Captcha(c *gin.Context) {
	driver := base64Captcha.NewDriverDigit(global.Config.Captcha.ImgHeight,
		global.Config.Captcha.ImgWidth,
		global.Config.Captcha.KeyLong,
		global.Config.Captcha.MaxSkew,
		global.Config.Captcha.DotCount)

	cp := base64Captcha.NewCaptcha(driver, store.UseWithCtx(c))

	id, b64s, err := cp.Generate()
	if err != nil {
		global.Log.Error("Captcha fail", zap.Error(err))

		response.ErrorInternal(c)
		return
	}

	response.SuccessWithData(c, resp.CaptchaResponse{
		CaptchaID:         id,
		CaptchaImgContent: b64s,
	})
}

func (a *LoginApi) Login(c *gin.Context) {
	var r request.Login
	msg := utils.Verify(c, &r)
	if msg != "" {
		response.ValidateFail(c, msg)
		return
	}

	// 验证码
	ok := store.Verify(r.CaptchaID, r.CaptchaCode, true)
	if !ok {
		response.CaptchaMistake(c)
		return
	}

	data, err := loginService.Login(r)
	if err != nil {
		response.DbError(c)
		return
	}

	response.SuccessWithData(c, data)
}

func (a *LoginApi) Refresh(c *gin.Context) {
	
}
