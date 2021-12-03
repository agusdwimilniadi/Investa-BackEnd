package routes

import (
	bankController "investaBackend/controllers/bank"
	investasi "investaBackend/controllers/investasi"
	sektor "investaBackend/controllers/sektor"

	ProyekMitraController "investaBackend/controllers/proyek_mitra"
	userInvestasiController "investaBackend/controllers/user_investasi"
	userController "investaBackend/controllers/users"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

type RouteControllerList struct {
	UserController         userController.UserController
	BankController         bankController.BankController
	ProyekMitraController  ProyekMitraController.ProyekMitraController
	UserInvestorController userInvestasiController.UserInvestasiController
	InvestasiController    investasi.InvestasiController
	SektorController       sektor.SektorController
	JWTConfig              middleware.JWTConfig
}

func (controller RouteControllerList) RouteRegister(c *echo.Echo) {
	// USER PROYEK
	users := c.Group("/user")
	users.GET("/masuk", controller.UserController.Login)
	users.POST("/daftar", controller.UserController.Register)

	// USER INVESTOR
	usersInvestor := c.Group("/investor")
	usersInvestor.GET("/masuk", controller.UserInvestorController.Login)
	usersInvestor.POST("/daftar", controller.UserInvestorController.Register)

	// Proyek
	c.GET("/proyek", controller.ProyekMitraController.GetAllDataProyek)
	c.POST("/proyek", controller.ProyekMitraController.CreateProyekController, middleware.JWTWithConfig(controller.JWTConfig))
	c.PUT("/proyek/:id", controller.ProyekMitraController.UpdateProyekController, middleware.JWTWithConfig(controller.JWTConfig))
	c.GET("/proyek/:id", controller.ProyekMitraController.GetAllDataByIdController)
	c.DELETE("/proyek/:id", controller.ProyekMitraController.DeleteProyekController, middleware.JWTWithConfig(controller.JWTConfig))
	c.GET("/proyek/sektor/:id", controller.ProyekMitraController.GetAllDataBySektorController)
	// c.GET("/proyek/total/:id", controller.ProyekMitraController.GetAllDataBySektorController)

	c.GET("/bank", controller.BankController.GetAllDataBank, middleware.JWTWithConfig(controller.JWTConfig))
	c.POST("/bank", controller.BankController.InsertBank, middleware.JWTWithConfig(controller.JWTConfig))

	c.POST("/investasi", controller.InvestasiController.InsertInvestasi, middleware.JWTWithConfig(controller.JWTConfig))
	c.GET("/investasi/:id", controller.InvestasiController.TotalInvestasiByIdController)

	c.POST("sektor", controller.SektorController.InsertSektor)

}
