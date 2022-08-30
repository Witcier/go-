package response

import (
	"witcier/go-api/model"
)

type CaptchaResponse struct {
	CaptchaID         string `json:"captchaID"`
	CaptchaImgContent string `json:"captchaImgContent"`
}

type LoginResponse struct {
	User      model.User `json:"user"`
	Token     string     `json:"token"`
	Type      string     `json:"type"`
	ExpiresAt int64      `json:"expiredIn"`
}
