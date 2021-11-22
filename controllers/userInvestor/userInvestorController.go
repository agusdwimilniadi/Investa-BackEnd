package userinvestor

import (
	"investaBackend/configs"
	"investaBackend/middlewares"
	"investaBackend/models"
	"investaBackend/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func RegisterUserInvestor(e echo.Context) error {
	var userInvestor models.UserInvestasi

	e.Bind(&userInvestor)
	configs.DB.Create(&userInvestor)
	if userInvestor.Password == "" {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Kosong",
			Data:    nil,
		})
	}

	return e.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    userInvestor,
	})
}

func LoginUserInvestor(e echo.Context) error {
	var userInvestor models.UserInvestasi

	e.Bind(&userInvestor)

	// err := configs.DB.Where("email = ? AND password = ?", userInvestor.Email, userInvestor.Password).First(&userInvestor)
	if err := configs.DB.Where("email = ? AND password = ?", userInvestor.Email, userInvestor.Password).First(&userInvestor).Error; err != nil {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Email Password salah",
			Data:    nil,
		})
	}
	if userInvestor.Password == "" {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Password Kosong",
			Data:    nil,
		})
	}

	// GENERATE TOKEN
	token, errors := middlewares.GenereteTokenJWT(int(userInvestor.ID))

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
			"Role":         "User Investor",
			"Nama Lengkap": userInvestor.Name,
			"Email":        userInvestor.Email,
			"token":        token,
		},
	})
}
