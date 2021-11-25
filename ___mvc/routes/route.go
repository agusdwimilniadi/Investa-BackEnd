package routes

import (
	"investaBackend/constant"
	bank "investaBackend/controllers/bank"
	investasi "investaBackend/controllers/investasi"
	loginInvestasi "investaBackend/controllers/userInvestor"
	loginProyek "investaBackend/controllers/userProyek"

	middlewares "investaBackend/middlewares"

	echo "github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
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

	// e.POST("/investasi", investasi.Investasi)
	eInvestasi := e.Group("/investasi")
	config := middleware.JWTConfig{
		Claims:     &middlewares.JwtCustomClaims{},
		SigningKey: []byte(constant.SECRET_JWT),
	}
	eInvestasi.Use(middleware.JWTWithConfig(config))
	eInvestasi.POST("/", investasi.Investasi)
	return e
}
