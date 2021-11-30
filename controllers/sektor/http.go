package sektor

import (
	"investaBackend/business/sektor"
	"investaBackend/controllers"
	"investaBackend/controllers/sektor/request"
	"investaBackend/controllers/sektor/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type SektorController struct {
	usecase sektor.UseCase
}

func NewSektorController(sc sektor.UseCase) *SektorController {
	return &SektorController{
		usecase: sc,
	}
}

func (controller SektorController) InsertSektor(c echo.Context) error {
	ctx := c.Request().Context()
	var insertSektor request.InsertSektor

	err := c.Bind(&insertSektor)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "error bind", err)
	}
	sektors, err := controller.usecase.InsertSektor(insertSektor.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "data kosong", err)
	}
	return controllers.SuccessResponse(c, response.FromDomain(sektors))

}
