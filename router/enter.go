package router

type RouterGroup struct {
	UserRouter
	LoginRouter
	PositionRouter
	DepartmentRouter
	MenuRouter
}

var RouterGroupApp = new(RouterGroup)
