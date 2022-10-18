package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type PermissionRouter struct{}

func (r *MenuRouter) InitPermissionRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	permissionApi := v1.ApiGroup.PermissionApi
	{
		router.GET("/permissions", permissionApi.List)
		router.POST("/permissions", permissionApi.Store)
		router.PATCH("/permissions/:id", permissionApi.Update)
		router.DELETE("/permissions/:id", permissionApi.Delete)
	}
}
