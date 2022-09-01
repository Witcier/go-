package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type DepartmentApi struct{}

func (a *DepartmentApi) ListDepartment(c *gin.Context) {
	var r request.List
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	data, err := departmentService.ListDepartment(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.SuccessWithData(c, data)
}

func (a *DepartmentApi) StoreDepartment(c *gin.Context) {
	var r request.StoreDepartment
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := departmentService.StoreDepartment(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *DepartmentApi) UpdateDepartment(c *gin.Context) {
	var r request.UpdateDepartment
	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := departmentService.UpdateDepartment(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *DepartmentApi) DeleteDepartment(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	if err := departmentService.DeleteDepartment(id); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}
