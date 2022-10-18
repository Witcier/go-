package service

type Enter struct {
	LoginService
	UserService
	DepartmentService
	PositionService
	MenuService
	PersonService
}

var ServiceGroup = new(Enter)
