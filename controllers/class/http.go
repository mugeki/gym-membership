package class

import (
	"gym-membership/business/class"
	controller "gym-membership/controllers"
	"gym-membership/controllers/class/request"
	"gym-membership/controllers/class/response"
	"net/http"
	"strconv"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ClassController struct {
	classUsecase class.Usecase
}

func NewClassController(Usecase class.Usecase) *ClassController {
	return &ClassController{
		classUsecase: Usecase,
	}
}

func (ctrl *ClassController) Insert(c echo.Context) error {
	req := request.Class{}
	domain := class.Domain{}
	if err := c.Bind(&req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	if _, err := govalidator.ValidateStruct(req); err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	copier.Copy(&domain, &req)
	data, err := ctrl.classUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusConflict, err)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, data)
}

func (ctrl *ClassController) GetAll(c echo.Context) error {
	title := c.QueryParam("name")
	page, _ := strconv.Atoi(c.QueryParam("page"))
	res := []response.Class{}
	if page <= 0 {
		page = 1
	}
	data, offset, limit, totalData, err := ctrl.classUsecase.GetAll(title, page)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	resPage := response.Page{
		Limit:     limit,
		Offset:    offset,
		TotalData: totalData,
	}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, data)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, res, resPage)
}

func (ctrl *ClassController) UpdateClassByID(c echo.Context) error {
	// println("cek param path", c.QueryParam("id"))
	req := request.ClassUpdate{}
	res := response.Class{}
	domain := class.Domain{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	articleId, _ := strconv.Atoi(c.Param("idClass"))
	copier.Copy(&domain, &req)
	println("controller 1", req.Name)
	data, err := ctrl.classUsecase.UpdateClassByID(uint(articleId), &domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}
