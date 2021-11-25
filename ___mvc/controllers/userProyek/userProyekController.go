package loginproyek

import (
	"investaBackend/configs"
	"investaBackend/middlewares"
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

func LoginUserProyek(e echo.Context) error {
	var userProyek userProyek.UserProyek

	e.Bind(&userProyek)

	// err := configs.DB.Where("email = ? AND password = ?", userProyek.Email, userProyek.Password).First(&userProyek)
	if err := configs.DB.Where("email = ? AND password = ?", userProyek.Email, userProyek.Password).First(&userProyek).Error; err != nil {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email Password salah",
			Data:    nil,
		})
	}
	if userProyek.Password == "" {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Kosong",
			Data:    nil,
		})
	}

	// GENERATE TOKEN
	token, errors := middlewares.GenereteTokenJWT(int(userProyek.ID))

	if errors != nil {
		return e.JSON(http.StatusInternalServerError, response.BaseResponse{
			Code:    http.StatusInternalServerError,
			Message: "Error generate JWT",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Succes",
		Data: map[string]string{
			"Role":         "User Proyek",
			"Nama Lengkap": userProyek.Name,
			"Email":        userProyek.Email,
			"token":        token,
		},
	})
}
