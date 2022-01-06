package classifications

import (
	"gym-membership/business/classification"
	controller "gym-membership/controllers"
	"gym-membership/controllers/classifications/request"
	"gym-membership/controllers/classifications/response"

	// "gym-membership/controllers/articles/response"
	"net/http"

	"github.com/asaskevich/govalidator"
	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type ClassificationController struct {
	classificationUsecase classification.Usecase
}

func NewClassificationController(Usecase classification.Usecase) *ClassificationController {
	return &ClassificationController{
		classificationUsecase: Usecase,
	}
}

func (ctrl *ClassificationController) GetAll(c echo.Context) error {
	data, err := ctrl.classificationUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := []response.Classification{}
	copier.Copy(&res, &data)
	if len(data) == 0 {
		return controller.NewSuccessResponse(c, http.StatusNoContent, data)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *ClassificationController) Insert(c echo.Context) error {
	req := request.Classification{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	domain := classification.Domain{}
	copier.Copy(&domain,&req)
	data, err := ctrl.classificationUsecase.Insert(&domain)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res :=response.Classification{}
	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}
