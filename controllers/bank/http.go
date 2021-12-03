package bank

import (
	"investaBackend/business/bank"
	"investaBackend/controllers"
	"investaBackend/controllers/bank/request"
	"investaBackend/controllers/bank/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

type BankController struct {
	usecase bank.UseCase
}

func NewBankController(bc bank.UseCase) *BankController {
	return &BankController{
		usecase: bc,
	}
}

func (controller BankController) InsertBank(c echo.Context) error {
	ctx := c.Request().Context()
	var insertBank request.InsertBank
	err := c.Bind(&insertBank)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error bind", err)
	}
	banks, err := controller.usecase.InsertBank(insertBank.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Data bank kosong", err)
	}

	return controllers.SuccessResponse(c, response.FromDomain(banks))
}

func (controller BankController) GetAllDataBank(c echo.Context) error {
	ctx := c.Request().Context()
	getData, err := controller.usecase.GetAllDataBank(ctx)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Fail", err)
	}
	var responses []response.InsertBankResponse
	for _, val := range getData {
		responses = append(responses, response.FromDomain(val))
	}
	return controllers.SuccessResponse(c, map[string]interface{}{
		"proyek": responses,
	})
	// return controllers.SuccessResponse(c, getData)
}
