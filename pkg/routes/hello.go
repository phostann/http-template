package routes

import (
	"github.com/labstack/echo/v4"
)

func (r *Route) helloRoutes(group *echo.Group) {
	router := group.Group("/hello")
	router.GET("", r.ctrl.Hello)
}
