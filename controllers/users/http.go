package users

import (
	"investaBackend/business/users"
	"investaBackend/controllers"
	"investaBackend/controllers/users/request"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserController struct {
	usecase users.UserUseCaseInterface
}

func NewUserController(uc users.UserUseCaseInterface) *UserController {
	return &UserController{
		usecase: uc,
	}

}

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
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

func (ctrl *UserController) Register(c echo.Context) error {
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
