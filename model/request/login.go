package request

type Login struct {
	Username    string `json:"username" binding:"required"`
	Password    string `json:"password" binding:"required"`
	CaptchaID   string `json:"captchaID" binding:"required"`
	CaptchaCode string `json:"captchaCode" binding:"required"`
}
