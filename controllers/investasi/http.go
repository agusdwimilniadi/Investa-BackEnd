package investasi

import (
	"investaBackend/business/investasi"
	"investaBackend/controllers"
	"investaBackend/controllers/investasi/request"
	"investaBackend/controllers/investasi/response"
	"strconv"

	"net/http"

	"github.com/labstack/echo/v4"
)

type InvestasiController struct {
	usecase investasi.Usecase
}

func NewInvestasiController(investUseCasee investasi.Usecase) *InvestasiController {
	return &InvestasiController{
		usecase: investUseCasee,
	}
}

func (control InvestasiController) InsertInvestasi(c echo.Context) error {
	ctx := c.Request().Context()
	var insertInvestasi request.InsertInvestasiReq

	err := c.Bind(&insertInvestasi)

	if err != nil {
		controllers.ErrorResponse(c, http.StatusInternalServerError, "Error, data harus lengkap", err)

	}
	invest, err := control.usecase.InsertInvestasi(insertInvestasi.ToDomainReq(), ctx)
	return controllers.SuccessResponse(c, response.FromDomain(invest))
}

func (handler InvestasiController) TotalInvestasiByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	total, err := handler.usecase.TotalInvestasiByIdController(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}
	return controllers.SuccessResponse(c, total)
}
