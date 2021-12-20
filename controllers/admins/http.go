package admins

import (
	"gym-membership/business/admins"
	controller "gym-membership/controllers"
	"gym-membership/controllers/admins/request"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/labstack/echo/v4"
)

type AdminController struct {
	adminUsecase admins.Usecase
}

func NewAdminController(Usecase admins.Usecase) *AdminController {
	return &AdminController{
		adminUsecase: Usecase,
	}
}

func (ctrl *AdminController) Register(c echo.Context) error {
	req := request.Admins{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.adminUsecase.Register(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}
	// println(req, "ff")

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *AdminController) Login(c echo.Context) error {
	req := request.AdminsLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	token, err := ctrl.adminUsecase.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	res := struct {
		Token string `json:"token"`
	}{Token: token}

	return controller.NewSuccessResponse(c, res)
}
