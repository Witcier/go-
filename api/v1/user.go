package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (a *UserApi) ListUser(c *gin.Context) {
	var r request.List
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	data, err := userService.ListUser(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.SuccessWithData(c, data)
}

func (a *UserApi) StoreUser(c *gin.Context) {
	var r request.StoreUser
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := userService.StoreUser(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *UserApi) UpdateUser(c *gin.Context) {
	var r request.UpdateUser
	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := userService.UpdateUser(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *UserApi) DeleteUser(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	if err := userService.DeleteUser(id); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *UserApi) Role(c *gin.Context) {
	var r request.StoreUserRole

	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := userService.Role(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}
