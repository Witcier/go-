package v1

import "witcier/go-api/service"

type Enter struct {
	UserApi
}

var ApiGroup = new(Enter)

var (
	userService = service.ServiceGroup.UserService
)
