package routes

import (
	bank "investaBackend/controllers/bank"
	loginInvestasi "investaBackend/controllers/userInvestasi"
	loginProyek "investaBackend/controllers/userProyek"

	middlewares "investaBackend/middlewares"

	echo "github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	middlewares.LogMiddleware(e)
	e.POST("/add/bank", bank.InsertBank)

	e.POST("/register/proyek", loginProyek.RegisterUserProyek)
	e.POST("/register/investasi", loginInvestasi.RegiterUserInvestasi)
	return e
}
