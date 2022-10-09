package routes

import (
	"github.com/labstack/echo/v4"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func (r *Route) swaggerRoute(group *echo.Group) {
	router := group.Group("/swagger")
	router.GET("*", echoSwagger.WrapHandler)
}
