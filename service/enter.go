package service

type Enter struct {
	LoginService
	UserService
	DepartmentService
	PositionService
	MenuService
	PersonService
	RoleService
}

var ServiceGroup = new(Enter)
