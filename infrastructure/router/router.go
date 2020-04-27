
package router

import (
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"github.com/linda/auth/interface/controller"
)

func NewRouter(e *echo.Echo, c controller.AppController) *echo.Echo {
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.POST("/signup", func(context echo.Context) error { return c.SignIn(context) })

	return e
}