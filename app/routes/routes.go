package routes

import (
	bankController "investaBackend/controllers/bank"
	userInvestasiController "investaBackend/controllers/user_investasi"
	userController "investaBackend/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController         userController.UserController
	BankController         bankController.BankController
	UserInvestorController userInvestasiController.UserInvestasiController
	JWTConfig              middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.POST("/login", controller.UserController.Login)

	usersInvestor := c.Group("/investor")
	usersInvestor.POST("/login", controller.UserInvestorController.Login)

	c.POST("/bank", controller.BankController.InsertBank)
}
