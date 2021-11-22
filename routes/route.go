package routes

import (
	bank "investaBackend/controllers/bank"
	loginInvestasi "investaBackend/controllers/userInvestor"
	loginProyek "investaBackend/controllers/userProyek"

	middlewares "investaBackend/middlewares"

	echo "github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)
	e.POST("/add/bank", bank.InsertBank)

	eLogin := e.Group("/login")
	eLogin.GET("/proyek", loginProyek.LoginUserProyek)
	eLogin.GET("/investasi", loginInvestasi.LoginUserInvestor)

	eRegister := e.Group("/register")

	eRegister.POST("/proyek", loginProyek.RegisterUserProyek)
	eRegister.POST("/investasi", loginInvestasi.RegisterUserInvestor)
	return e
}
