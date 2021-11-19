package routes

import (
	controllers "investaBackend/controllers"
	middlewares "investaBackend/middlewares"

	echo "github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)
	e.GET("/", controllers.HelloController)
	return e
}
