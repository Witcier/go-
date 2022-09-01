package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type PositionApi struct{}

func (a *PositionApi) ListPosition(c *gin.Context) {
	data, err := positionService.ListPosition()
	if err != nil {
		utils.DbError(c)
	}

	utils.SuccessWithData(c, data)
}

func (a *PositionApi) StorePosition(c *gin.Context) {
	var r request.StorePosition
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := positionService.StorePosition(r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *PositionApi) UpdatePosition(c *gin.Context) {
	var r request.UpdatePosition
	id := utils.ParseParamID(c, "id")
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := positionService.UpdatePosition(id, r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *PositionApi) DeletePosition(c *gin.Context) {
	id := utils.ParseParamID(c, "id")

	if err := positionService.DeletePosition(id); err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}
