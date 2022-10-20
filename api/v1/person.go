package v1

import (
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type PersonApi struct{}

func (a *PersonApi) PersonDetail(c *gin.Context) {
	user, err := personService.PersonDetail(c)
	if err != nil {
		utils.Unauthorized(c)
		return
	}

	utils.SuccessWithData(c, user)
}

func (a *PersonApi) PersonUpdatePassword(c *gin.Context) {
	var r request.PersonUpdatePassword
	errMsg := utils.Verify(c, &r)
	if errMsg != "" {
		utils.ValidateFail(c, errMsg)
		return
	}

	err := personService.PersonUpdatePassword(c, r)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.Success(c)
}

func (a *PersonApi) Menu(c *gin.Context) {
	data, err := personService.Menu(c)
	if err != nil {
		utils.DbError(c)
		return
	}

	utils.SuccessWithData(c, data)
}
