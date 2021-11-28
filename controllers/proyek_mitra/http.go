package proyek_mitra

import (
	"investaBackend/business/proyek_mitra"
	"investaBackend/controllers"
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
)

type ProyekMitraController struct {
	ProyekControllerUseCase proyek_mitra.Usecase
}

func NewProyekMitraController(proyekMitraUsecase proyek_mitra.Usecase) *ProyekMitraController {
	return &ProyekMitraController{
		ProyekControllerUseCase: proyekMitraUsecase,
	}
}

func (handler ProyekMitraController) GetAllDataProyek(c echo.Context) error {
	ctx := c.Request().Context()
	proyekMitra, err := handler.ProyekControllerUseCase.GetAllDataProyek(ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Failed fetch data", err)
	}
	return controllers.SuccessResponse(c, proyekMitra)
}

func (handler ProyekMitraController) GetAllDataByIdController(c echo.Context) error {
	id, _ := strconv.Atoi(c.Param("id"))
	ctx := c.Request().Context()

	proyekMitra, err := handler.ProyekControllerUseCase.GetAllDataByIdController(ctx, id)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}
	return controllers.SuccessResponse(c, proyekMitra)
}

func (handler ProyekMitraController) CreateProyekController(c echo.Context) error {
	proyekMitraInsert := proyek_mitra.Domain{}
	c.Bind(&proyekMitraInsert)
	ctx := c.Request().Context()

	proyekMitra, err := handler.ProyekControllerUseCase.CreateProyekController(ctx, proyekMitraInsert)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, err.Error(), err)
	}
	return controllers.SuccessResponse(c, proyekMitra)
}
