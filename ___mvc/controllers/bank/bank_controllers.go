package bank

import (
	"investaBackend/configs"
	"investaBackend/models"
	"investaBackend/models/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

func InsertBank(e echo.Context) error {
	var bank models.BankCode
	e.Bind(&bank)

	result := configs.DB.Create(&bank)

	if bank.KodeBank == "" || bank.NamaBank == "" {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data tidak boleh kosong",
			Data:    "No data",
		})
	}

	if result.Error != nil {
		return e.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Failed Created Data",
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    bank,
	})
}
