package middleware

import (
	"witcier/go-api/model/common/response"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetBearerToken(c)
		if token == "" {
			response.Unauthorized(c)
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil {
			response.Unauthorized(c)
			return
		}

		c.Set("claims", claims)
		c.Next()
	}
}
