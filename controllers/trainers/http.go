package trainers

import (
	"gym-membership/business/trainers"
	controller "gym-membership/controllers"
	"gym-membership/controllers/trainers/response"
	"net/http"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
)

type TrainerController struct {
	trainerUsecase trainers.Usecase
}

func NewTrainerController(Usecase trainers.Usecase) *TrainerController {
	return &TrainerController{
		trainerUsecase: Usecase,
	}
}

func (ctrl *TrainerController) GetAll(c echo.Context) error {
	data, err := ctrl.trainerUsecase.GetAll()
	res := []response.Trainers{}
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}
