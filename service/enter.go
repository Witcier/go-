package service

type Enter struct {
	LoginService
	UserService
	DepartmentService
}

var ServiceGroup = new(Enter)
