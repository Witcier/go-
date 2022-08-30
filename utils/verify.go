package utils

import (
	"witcier/go-api/global"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func Verify(c *gin.Context, request interface{}) string {
	if err := c.ShouldBindJSON(request); err != nil {
		if errors, ok := err.(validator.ValidationErrors); !ok {
			return err.Error()
		} else {
			validMsg := errors.Translate(global.Trans)
			for _, v := range validMsg {
				return v
			}
		}
	}

	return ""
}
