package users

import (
	"investaBackend/business/users"
	"investaBackend/controllers"
	"investaBackend/controllers/users/request"
	"investaBackend/controllers/users/response"

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

func (controller UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()
	var userLogin request.UserLogin
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
	return controllers.SuccessResponse(c, response.FromDomain(user))
}

func (controller UserController) GetAllUsers(c echo.Context) error {
	return controllers.SuccessResponse(c, response.UserResponse{})
}
