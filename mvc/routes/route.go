package routes

import (
	"investaBackend/controllers/proyekMitra"

	"github.com/labstack/echo/v4"
)

func New() *echo.Echo {
	e := echo.New()

	e.GET("/proyek", proyekMitra.GetAllDataProyek)
}
