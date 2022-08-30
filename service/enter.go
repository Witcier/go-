package service

type Enter struct {
	LoginService
	UserService
}

var ServiceGroup = new(Enter)
