package routes

import (
	"github.com/labstack/echo/v4"

	"github.com/phostann/template/http/app/controllers"
	"github.com/phostann/template/http/pkg/log"
)

type Route struct {
	ctrl   *controllers.Controller
	logger *log.Logger
}

func NewRoute(ctrl *controllers.Controller, logger *log.Logger) *Route {
	return &Route{
		ctrl:   ctrl,
		logger: logger,
	}
}

func (r *Route) SetUp(router *echo.Group) {
	r.helloRoutes(router)
	r.swaggerRoute(router)
}
