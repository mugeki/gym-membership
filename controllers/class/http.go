package class

import (
	"gym-membership/business/class"
	controller "gym-membership/controllers"
	"gym-membership/controllers/class/request"
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

	return controller.NewSuccessResponse(c, data)
}

func (ctrl *ClassController) UpdateKuota(c echo.Context) error {
	// req := request.Videos{}
	// err := c.Bind(&req)
	// if err != nil {
	// 	return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	// }

	// _, err = govalidator.ValidateStruct(req)
	// if err != nil {
	// 	return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	// }

	classId, _ := strconv.Atoi(c.Param("idClass"))
	// domain := videos.Domain{}
	// copier.Copy(&domain, &req)
	stringStatus, err := ctrl.classUsecase.UpdateKuota(int(classId))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, stringStatus)
}
