package userinvestasi

import (
	"investaBackend/configs"
	userInvestasi "investaBackend/models"
	"investaBackend/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegiterUserInvestasi(e echo.Context) error {
	var userInvestasi userInvestasi.UserProyek

	e.Bind(&userInvestasi)
	configs.DB.Create(&userInvestasi)
	if userInvestasi.Password == "" {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Kosong",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userInvestasi,
	})
}
