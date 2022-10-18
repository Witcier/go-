package service

type Enter struct {
	LoginService
	UserService
	DepartmentService
	PositionService
	MenuService
	PersonService
	RoleService
	PermissionService
}

var ServiceGroup = new(Enter)
