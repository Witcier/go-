package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type MenuApi struct{}

func (a *MenuApi) ListMenu(c *gin.Context) {
	data, err := menuService.ListMenu()
	if err != nil {
		utils.DbError(c)
	}

	utils.SuccessWithData(c, data)
}

func (a *MenuApi) StoreMenu(c *gin.Context) {
	var r request.StoreMenu
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := menuService.StoreMenu(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *MenuApi) UpdateMenu(c *gin.Context) {
	var r request.UpdateMenu
	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	if err := menuService.UpdateMenu(id, r); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *MenuApi) DeleteMenu(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	if err := menuService.DeleteMenu(id); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}
