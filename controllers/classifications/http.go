package classifications

import (
	"gym-membership/business/classification"
	controller "gym-membership/controllers"
	"gym-membership/controllers/classifications/request"

	// "gym-membership/controllers/articles/response"
	"net/http"

	"github.com/asaskevich/govalidator"
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

func (ctrl *ClassificationController) Insert(c echo.Context) error {
	// println("cek param path", c.QueryParam("id"))
	req := request.Classification{}
	err := c.Bind(&req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	_, err = govalidator.ValidateStruct(req)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusBadRequest, err)
	}

	// adminId := 1 //temporary adminID
	data, err := ctrl.classificationUsecase.Insert(req.ToDomain())
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, data)
}
