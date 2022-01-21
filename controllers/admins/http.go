package admins

import (
	"gym-membership/app/middleware"
	"gym-membership/business/admins"
	controller "gym-membership/controllers"
	"gym-membership/controllers/admins/request"
	"gym-membership/controllers/admins/response"
	"net/http"
	"strconv"

	"github.com/jinzhu/copier"

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
	domain := admins.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}
	copier.Copy(&domain, &req)
	data, err := ctrl.adminUsecase.Register(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}
	res := response.Admins{}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *AdminController) Login(c echo.Context) error {
	req := request.AdminsLogin{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	data, err := ctrl.adminUsecase.Login(req.Username, req.Password)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusUnauthorized, err)
	}
	res := response.Admins{}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *AdminController) Update(c echo.Context) error {
	req := request.Admins{}
	domain := admins.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	userId := uint(middleware.GetUser(c).ID)
	data, err := ctrl.adminUsecase.Update(userId, &domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := response.Admins{}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *AdminController) GetAll(c echo.Context) error {
	name := c.QueryParam("name")
	id, _ := strconv.Atoi(c.QueryParam("id"))
	page, _ := strconv.Atoi(c.QueryParam("page"))
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.adminUsecase.GetAll(uint(id), name, page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := []response.Admins{}
	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, res)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
}