package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type UserRouter struct{}

func (r *UserRouter) InitUserRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	userApi := v1.ApiGroup.UserApi
	{
		router.POST("/users", userApi.StoreUser)
		router.PATCH("/users/:id", userApi.UpdateUser)
	}
}
