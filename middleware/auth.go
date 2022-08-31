package middleware

import (
	"time"
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
		if err != nil || claims.ExpiresAt.Unix() < time.Now().Unix() {
			response.Unauthorized(c)
			return
		}

		c.Next()
	}
}
