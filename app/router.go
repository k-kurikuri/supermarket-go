package app

import (
	"github.com/labstack/echo"
)

type Router struct {
}

func NewRouter() *Router {
	return &Router{}
}

func (r *Router) Boot(e *echo.Echo) {
	control := NewController()

	e.GET("/", control.Index)
	e.POST("/add", control.AddTask)
	e.DELETE("/delete", control.DeleteTask)
	e.PUT("/update", control.PutTask)
}
