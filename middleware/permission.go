package middleware

import (
	"witcier/go-api/model"
	"witcier/go-api/utils"

	"github.com/gin-gonic/gin"
)

type CheckPermission struct {
	User        *model.User
	Path        string
	Method      string
	Permissions []model.Permission
}

func Permission() gin.HandlerFunc {
	return func(c *gin.Context) {
		user, err := utils.GetAuthUser(c)
		if err != err {
			utils.Unauthorized(c)
			return
		}

		// 判断是否管理员
		if ok := CheckUserIsAdmin(user); ok {
			c.Next()
			return
		}

		permissions, err := utils.GetUserPermission(user)
		if err != nil {
			utils.Denied(c)
			return
		}

		checkPermission := CheckPermission{
			Path:        c.Request.URL.Path,
			Method:      c.Request.Method,
			Permissions: permissions,
		}

		if ok := checkPermission.CheckUserPermission(); !ok {
			utils.Denied(c)
		}

		c.Next()
	}
}

func CheckUserIsAdmin(user model.User) bool {
	return user.IsAdmin
}

func (m *CheckPermission) CheckUserPermission() bool {
	return false
}