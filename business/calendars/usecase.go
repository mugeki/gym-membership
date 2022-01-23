package calendars

import (
	"fmt"
	"gym-membership/business"
)

type calendarUsecase struct {
	calendarRepository Repository
	// jwtAuth            *middleware.ConfigJWT
}

func NewCalendarUsecase(calendarRepo Repository) Usecase {
	return &calendarUsecase{
		calendarRepository: calendarRepo,
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

func (uc *calendarUsecase) GetAll() (Event, error) {
	fmt.Println("=================== bussineslayer")
	data, err := uc.calendarRepository.GetAll()
	if err != nil {
		return Event{}, business.ErrInternalServer
	}
	return data, nil
}
