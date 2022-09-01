package v1

import "witcier/go-api/service"

type Enter struct {
	UserApi
	LoginApi
	DepartmentApi
	PositionApi
}

var ApiGroup = new(Enter)

var (
	userService       = service.ServiceGroup.UserService
	loginService      = service.ServiceGroup.LoginService
	departmentService = service.ServiceGroup.DepartmentService
	positionService   = service.ServiceGroup.PositionService
)
