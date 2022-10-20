package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type RoleApi struct{}

func (a *RoleApi) List(c *gin.Context) {
	var r request.List
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	data, err := roleService.List(r)
	if err != nil {
		utils.DbError(c)
	}

	utils.SuccessWithData(c, data)
}

func (a *RoleApi) Store(c *gin.Context) {
	var r request.StoreRole
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := roleService.Store(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *RoleApi) Update(c *gin.Context) {
	var r request.UpdateRole
	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := roleService.Update(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *RoleApi) Delete(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	if err := roleService.Delete(id); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *RoleApi) Menu(c *gin.Context) {
	var r request.StoreRoleMenu
	id := utils.ParseParamID(c, "id")

	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := roleService.Menu(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *RoleApi) GetMenu(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	data, err := roleService.GetMenu(id)
	if err != nil {
		utils.DbError(c)
	}

	utils.SuccessWithData(c, data)
}

func (a *RoleApi) Permission(c *gin.Context) {
	var r request.StoreRolePermission
	id := utils.ParseParamID(c, "id")

	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := roleService.Permission(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *RoleApi) GetPermission(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	data, err := roleService.GetPermission(id)
	if err != nil {
		utils.DbError(c)
	}

	utils.SuccessWithData(c, data)
}
