package routes

import (
	bankController "investaBackend/controllers/bank"
	investasi "investaBackend/controllers/investasi"
	userInvestasiController "investaBackend/controllers/user_investasi"

	userController "investaBackend/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController         userController.UserController
	BankController         bankController.BankController
	UserInvestorController userInvestasiController.UserInvestasiController
	InvestasiController    investasi.InvestasiController

	JWTConfig middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	users := c.Group("/user")
	users.GET("/login", controller.UserController.Login)
	users.POST("/register", controller.UserController.Register)

	usersInvestor := c.Group("/investor")
	usersInvestor.POST("/login", controller.UserInvestorController.Login)

	c.POST("/bank", controller.BankController.InsertBank)

	c.POST("/investasi", controller.InvestasiController.InsertInvestasi)
}
