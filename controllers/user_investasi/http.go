package userinvestasi

import (
	"investaBackend/business/user_investasi"
	"investaBackend/controllers"
	"investaBackend/controllers/user_investasi/request"
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

func (controller *UserInvestasiController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserInvestorLogin
	c.Bind(&userLogin)

	token, err := controller.usecase.Login(ctx, userLogin.Email, userLogin.Password)
	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, "Email atau password salah", err)
	}
	responses := struct {
		Email string `json:"email"`
		Token string `json:"token"`
	}{
		Email: userLogin.Email,
		Token: token,
	}

	return controllers.SuccessResponse(c, responses)
}

func (ctrl *UserInvestasiController) Register(c echo.Context) error {
	ctx := c.Request().Context()

	req := request.UserRegister{}
	err := c.Bind(&req)

	if err != nil {
		return controllers.ErrorResponse(c, http.StatusBadRequest, " ", err)
	}
	err2 := ctrl.usecase.Register(ctx, req.ToDomain())
	if err2 != nil {
		return controllers.ErrorResponse(c, http.StatusInternalServerError, " ", err2)
	}

	return controllers.SuccessResponse(c, http.StatusOK)

}
