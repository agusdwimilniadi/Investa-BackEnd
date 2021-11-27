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

func (controller *UserController) Login(c echo.Context) error {
	ctx := c.Request().Context()

	// emails := c.Request().Header
	// password := c.Request().Header
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
	// return controllers.SuccessResponse(c,)
	// var userLogin request.UserLogin
	// err := c.Bind(&userLogin)

	// if err != nil {
	// 	return controllers.ErrorResponse(c, http.StatusInternalServerError, "Error bind", err)
	// }
	// // hashing, err := encrypt.Hash(userLogin.Password)
	// user, err := controller.usecase.Login(*userLogin.ToDomain(), ctx)
	// if err != nil {
	// 	return controllers.ErrorResponse(c, http.StatusNotFound, "Email password salah", err)
	// }
	// // fmt.Println(!encrypt.ValidateHash(userLogin.Password, hashing))
	// return controllers.SuccessResponse(c, response.FromDomain(user))
}

func (controller UserController) GetAllUsers(c echo.Context) error {
	return controllers.SuccessResponse(c, response.UserResponse{})
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
