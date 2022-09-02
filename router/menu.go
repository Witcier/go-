package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type MenuRouter struct{}

func (r *MenuRouter) InitMenuRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	menuApi := v1.ApiGroup.MenuApi
	{
		router.GET("/menus", menuApi.ListMenu)
		router.POST("/menus", menuApi.StoreMenu)
		router.PATCH("/menus/:id", menuApi.UpdateMenu)
		router.DELETE("/menus/:id", menuApi.DeleteMenu)
	}
}
