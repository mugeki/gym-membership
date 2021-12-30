package users

import (
	"gym-membership/business/users"
	controller "gym-membership/controllers"
	"gym-membership/controllers/users/request"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type UserController struct {
	userUsecase users.Usecase
}

func NewUserController(Usecase users.Usecase) *UserController {
	return &UserController{
		userUsecase: Usecase,
	}
}

func (ctrl *UserController) Register(c echo.Context) error {
	req := request.Users{}
	domain := users.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.userUsecase.Register(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

func (ctrl *UserController) Login(c echo.Context) error {
	req := request.UsersLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userUsecase.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, http.StatusOK, res)
}
