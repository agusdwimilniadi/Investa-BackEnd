package routes

import (
	bankController "investaBackend/controllers/bank"
	investasi "investaBackend/controllers/investasi"
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
	c.PUT("/proyek/:id", controller.ProyekMitraController.UpdateProyekController)
	c.GET("/proyek/:id", controller.ProyekMitraController.GetAllDataByIdController)
	c.DELETE("/proyek/:id", controller.ProyekMitraController.DeleteProyekController)
	c.GET("/proyek/sektor/:id", controller.ProyekMitraController.GetAllDataBySektorController)
	// c.GET("/proyek/total/:id", controller.ProyekMitraController.GetAllDataBySektorController)

	c.POST("/bank", controller.BankController.InsertBank)

	c.POST("/investasi", controller.InvestasiController.InsertInvestasi)
	c.GET("/investasi/:id", controller.InvestasiController.TotalInvestasiByIdController)

}
