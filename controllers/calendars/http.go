package calendars

import (
	"gym-membership/business/calendars"
	controller "gym-membership/controllers"
	"net/http"

	"github.com/labstack/echo/v4"
)

type CalendarsController struct {
	calendarsUsecase calendars.Usecase
}

func NewCalendarsController(Usecase calendars.Usecase) *CalendarsController {
	return &CalendarsController{
		calendarsUsecase: Usecase,
	}
}

func (ctrl *CalendarsController) GetAll(c echo.Context) error {
	
	data, err := ctrl.calendarsUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, data, nil)
}

func (ctrl *CalendarsController) AddGuest(c echo.Context) error {
	data, err := ctrl.calendarsUsecase.GetAll()
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, data, nil)
}
