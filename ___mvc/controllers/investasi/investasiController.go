package investasi

import (
	"investaBackend/configs"
	"investaBackend/models"
	"investaBackend/models/response"
	"net/http"

	"github.com/labstack/echo/v4"
)

func Investasi(e echo.Context) error {
	var investasi models.Investasi

	e.Bind(&investasi)
	configs.DB.Create(&investasi)
	if investasi.ProyekID == 0 || investasi.UserInvestasiID == 0 {
		return e.JSON(http.StatusBadRequest, response.BaseResponse{
			Code:    http.StatusBadRequest,
			Message: "Data proyek dan user tidak boleh kosong",
			Data:    nil,
		})
	}
	return e.JSON(http.StatusOK, response.BaseResponse{
		Code:    http.StatusOK,
		Message: "Success",
		Data:    investasi,
	})
}
