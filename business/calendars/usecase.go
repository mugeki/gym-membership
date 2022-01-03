package calendars

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
	"gym-membership/helper/encrypt"

	"github.com/google/uuid"
)

type calendarUsecase struct {
	calendarRepository 	Repository
	jwtAuth			*middleware.ConfigJWT
}

func NewCalendarUsecase(calendarRepo Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &calendarUsecase{
		calendarRepository: calendarRepo,
		jwtAuth: jwtauth,
	}
}

func 