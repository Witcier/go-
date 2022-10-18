package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type PermissionApi struct{}

func (a *PermissionApi) List(c *gin.Context) {
	data, err := permissionService.List()
	if err != nil {
		utils.DbError(c)
	}

	utils.SuccessWithData(c, data)
}

func (a *PermissionApi) Store(c *gin.Context) {
	var r request.StorePermission
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := permissionService.Store(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *PermissionApi) Update(c *gin.Context) {
	var r request.UpdatePermission
	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := permissionService.Update(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *PermissionApi) Delete(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	if err := permissionService.Delete(id); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}
