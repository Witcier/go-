package router

import (
	v1 "witcier/go-api/api/v1"

	"github.com/gin-gonic/gin"
)

type DepartmentRouter struct{}

func (r *LoginRouter) InitDepartmentRouter(Router *gin.RouterGroup) {
	router := Router.Group("")
	departmentApi := v1.ApiGroup.DepartmentApi
	{
		router.GET("/departments", departmentApi.ListDepartment)
		router.POST("/departments", departmentApi.StoreDepartment)
		router.PATCH("/departments/:id", departmentApi.UpdateDepartment)
		router.DELETE("/departments/:id", departmentApi.DeleteDepartment)
	}
}
