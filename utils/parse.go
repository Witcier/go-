package utils

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func ParseParamID(c *gin.Context, key string) uint {
	val := c.Param(key)
	id, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		return 0
	}
	return uint(id)
}
