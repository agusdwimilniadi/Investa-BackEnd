package loginproyek

import (
	"investaBackend/configs"
	userProyek "investaBackend/models"
	"investaBackend/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserProyek(e echo.Context) error {
	var userProyek userProyek.UserProyek

	e.Bind(&userProyek)
	configs.DB.Create(&userProyek)
	if userProyek.Password == "" {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Kosong",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userProyek,
	})
}
