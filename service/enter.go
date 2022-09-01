package service

type Enter struct {
	LoginService
	UserService
	DepartmentService
	PositionService
}

var ServiceGroup = new(Enter)
