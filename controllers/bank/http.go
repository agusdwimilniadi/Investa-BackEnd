package bank

import (
	"investaBackend/business/bank"
	"investaBackend/controllers"
	"investaBackend/controllers/bank/request"
	"investaBackend/controllers/bank/response"
	"log"
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
	log.Println("sebelum insert ke db")
	log.Println(insertBank)
	banks, err := controller.usecase.InsertBank(insertBank.ToDomain(), ctx)
	log.Println("setelah insert ke db")
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Data bank kosong", err)
	}

	return controllers.SuccessResponse(c, response.FromDomain(banks))
}
