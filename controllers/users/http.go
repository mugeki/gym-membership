package users

import (
	"gym-membership/business/users"
	controller "gym-membership/controllers"
	"gym-membership/controllers/users/request"
	"net/http"

	"github.com/asaskevich/govalidator"
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

func (ctrl *UserController) Register(c echo.Context) error{
	req := request.Users{}
	if err := c.Bind(&req); err != nil{
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.userUsecase.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *UserController) Login(c echo.Context) error{
	req := request.UsersLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.userUsecase.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	
	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c,res)
}
