package calendarsApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"gym-membership/business/calendars"
	"net/http"
)

type CalendarsApi struct {
	httpClient http.Client
}

func NewCalendarsApi() calendars.Repository {
	return &CalendarsApi{
		httpClient: http.Client{},
	}
}

func (ca *CalendarsApi) CreateEvent(eventData *calendars.Event) (calendars.Event, error) {
	calendarId := "qurrotaayunamiroh@gmail.com" //tempCalendarID

	requestLink := (fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/%s/events", calendarId))
	reqBody, err := json.Marshal(eventData)
	if err != nil {
		return calendars.Event{}, err
	}

	resp, err := http.NewRequest(http.MethodPut, requestLink, bytes.NewBuffer(reqBody))
	if err != nil {
		return calendars.Event{}, err
	}
	result := calendars.Event{}
	err = json.NewDecoder(resp.Body).Decode(&result)
	if err != nil {
		return calendars.Event{}, err
	}

	return result, nil

}

func (ca *CalendarsApi) AddGuest(eventId, emailGuest string) (calendars.Event, error) {
	calendarId := "qurrotaayunamiroh@gmail.com" //tempCalendarID

	requestLinkGetById := (fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/%s/events/%s", calendarId, eventId))

	resp, err := http.Get(requestLinkGetById)
	if err != nil {
		return calendars.Event{}, err
	}
	result := calendars.Event{}
	err = json.NewDecoder(resp.Body).Decode(&result)

	eventDataUpdate := result
	// eventDataUpdate
	requestLink := (fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/%s/events", calendarId))
	reqBody, err := json.Marshal(eventDataUpdate)
	if err != nil {
		return calendars.Event{}, err
	}

	respUpdate, errUpdate := http.NewRequest(http.MethodPut, requestLink, bytes.NewBuffer(reqBody))
	if errUpdate != nil {
		return calendars.Event{}, err
	}
	resultUpdate := calendars.Event{}
	err = json.NewDecoder(respUpdate.Body).Decode(&resultUpdate)
	if err != nil {
		return calendars.Event{}, err
	}

	return calendars.Event{}, nil

}
