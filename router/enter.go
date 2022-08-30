package router

type RouterGroup struct {
	UserRouter
	LoginRouter
}

var RouterGroupApp = new(RouterGroup)
