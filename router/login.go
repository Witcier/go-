package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type LoginRouter struct{}

func (r *LoginRouter) InitLoginRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	loginApi := v1.ApiGroup.LoginApi
	{
		router.GET("/captcha", loginApi.Captcha)
		router.POST("/login", loginApi.Login)
		router.PATCH("/refresh", loginApi.Refresh)
	}
}
