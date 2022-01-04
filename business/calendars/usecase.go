package calendars

import (
	"gym-membership/app/middleware"
	"gym-membership/business"
)

type calendarUsecase struct {
	calendarRepository Repository
	jwtAuth            *middleware.ConfigJWT
}

func NewCalendarUsecase(calendarRepo Repository, jwtauth *middleware.ConfigJWT) Usecase {
	return &calendarUsecase{
		calendarRepository: calendarRepo,
		jwtAuth:            jwtauth,
	}
}

func (uc *calendarUsecase) CreateEvent(EventData *Event) (Event, error) {

	data, err := uc.calendarRepository.CreateEvent(EventData)
	if err != nil {
		return Event{}, business.ErrInternalServer
	}
	return data, nil
}

func (uc *calendarUsecase) AddGuest(eventId, emailGuest string) (Event, error) {

	data, err := uc.calendarRepository.AddGuest(eventId, emailGuest)
	if err != nil {
		return Event{}, business.ErrInternalServer
	}
	return data, nil
}
