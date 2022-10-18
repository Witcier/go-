package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type PersonRouter struct{}

func (r *PersonRouter) InitPersonRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	personApi := v1.ApiGroup.PersonApi
	{
		router.GET("/person/detail", personApi.PersonDetail)
		router.PATCH("/person/update/password", personApi.PersonUpdatePassword)
	}
}
