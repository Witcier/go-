package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type RoleRouter struct{}

func (r *MenuRouter) InitRoleRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	roleApi := v1.ApiGroup.RoleApi
	{
		router.GET("/roles", roleApi.List)
		router.POST("/roles", roleApi.Store)
		router.PATCH("/roles/:id", roleApi.Update)
		router.DELETE("/roles/:id", roleApi.Delete)
	}
}
