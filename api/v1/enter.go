package v1

import "witcier/go-api/service"

type Enter struct {
	UserApi
	LoginApi
}

var ApiGroup = new(Enter)

var (
	userService  = service.ServiceGroup.UserService
	loginService = service.ServiceGroup.LoginService
)
