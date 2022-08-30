package utils

import (
	"strings"

	"github.com/gin-gonic/gin"
)

func GetBearerToken(c *gin.Context) string {
	token := c.Request.Header.Get("Authorization")
	// Should be a bearer token
	if len(token) > 6 && strings.ToUpper(token[0:7]) == "BEARER " {
		return token[7:]
	}
	return token
}
