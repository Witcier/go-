package router

type RouterGroup struct {
	UserRouter
	LoginRouter
	PositionRouter
	DepartmentRouter
	MenuRouter
	PersonRouter
}

var RouterGroupApp = new(RouterGroup)
