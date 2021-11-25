package userinvestasi

import (
	"investaBackend/business/user_investasi"
	"investaBackend/controllers"
	_request "investaBackend/controllers/user_investasi/request"
	_response "investaBackend/controllers/user_investasi/response"

	"net/http"

	"github.com/labstack/echo/v4"
)

type UserInvestasiController struct {
	usecase user_investasi.UserInvestasiUseCaseInterface
}

func NewUserInvestasiController(uc user_investasi.UserInvestasiUseCaseInterface) *UserInvestasiController {
	return &UserInvestasiController{
		usecase: uc,
	}
}

func (controller UserInvestasiController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin _request.UserInvestorLogin
	err := c.Bind(&userLogin)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error bind", err)
	}
	user, err := controller.usecase.Login(*userLogin.ToDomain(), ctx)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusNotFound, "Email password salah", err)
	}
	// if err != nil {

	// }
	return controllers.SuccessResponse(c, _response.FromDomain(user))
}
