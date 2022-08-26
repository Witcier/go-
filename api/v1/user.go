package v1

import (
	"fmt"
	"witcier/go-api/global"
	"witcier/go-api/model"
	"witcier/go-api/model/common/response"
	"witcier/go-api/model/request"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type UserApi struct{}

func (a *UserApi) StoreUser(c *gin.Context) {
	var r request.StoreUser
	if err1 := c.ShouldBindJSON(&r); err1 != nil {
		fmt.Println(err1)
	}
	fmt.Println(r)
	if err := utils.Verify(r); err != nil {
		response.ValidateFail(c)
		return
	}

	user := &model.User{
		Username:     r.Username,
		RealName:     r.RealName,
		Email:        r.Email,
		Phone:        r.Phone,
		Sex:          r.Sex,
		Birth:        r.Birth,
		EntryTime:    r.EntryTime,
		Status:       r.Status,
		IsAdmin:      r.IsAdmin,
		DepartmentID: r.DepartmentID,
		PositionID:   r.PositionID,
	}
	_, err := userService.StoreUser(*user)
	if err != nil {
		response.DbError(c)
		return
	}

	response.Success(c)
}

func (a *UserApi) UpdateUser(c *gin.Context) {
	var r request.UpdateUser
	id := utils.ParseParamID(c, "id")
	_ = c.ShouldBindJSON(&r)
	if err := utils.Verify(r); err != nil {
		response.ValidateFail(c)
		return
	}

	user := &model.User{
		ID: global.ID{
			ID: id,
		},
		Username:     r.Username,
		RealName:     r.RealName,
		Email:        r.Email,
		Phone:        r.Phone,
		Sex:          r.Sex,
		Birth:        r.Birth,
		EntryTime:    r.EntryTime,
		Status:       r.Status,
		IsAdmin:      r.IsAdmin,
		DepartmentID: r.DepartmentID,
		PositionID:   r.PositionID,
	}

	if err := userService.UpdateUser(*user); err != nil {
		response.DbError(c)
		return
	}

	response.Success(c)
}
