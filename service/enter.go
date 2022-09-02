package service

type Enter struct {
	LoginService
	UserService
	DepartmentService
	PositionService
	MenuService
}

var ServiceGroup = new(Enter)
