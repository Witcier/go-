package middleware

import (
	"fmt"
	"time"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

func Auth() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := utils.GetBearerToken(c)
		if token == "" {
			utils.Unauthorized(c)
			return
		}

		j := utils.NewJWT()

		claims, err := j.ParseToken(token)
		if err != nil || claims.ExpiresAt.Unix() < time.Now().Unix() {
			fmt.Println(err)
			utils.Unauthorized(c)
			return
		}

		c.Next()
	}
}
