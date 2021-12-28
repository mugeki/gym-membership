package trainers

import (
	"gym-membership/business/trainers"
	controller "gym-membership/controllers"
	"net/http"

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
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	// copier.Copy(&res, &data)
	return controller.NewSuccessResponse(c, http.StatusOK, data)
}
