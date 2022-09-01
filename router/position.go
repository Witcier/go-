package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type PositionRouter struct{}

func (r *LoginRouter) InitPositionRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	positionApi := v1.ApiGroup.PositionApi
	{
		router.GET("/positions", positionApi.ListPosition)
		router.POST("/positions", positionApi.StorePosition)
		router.PATCH("/positions/:id", positionApi.UpdatePosition)
		router.DELETE("/positions/:id", positionApi.DeletePosition)
	}
}
