package calendarsApi

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	// "os"

	"gym-membership/business/calendars"

	"google.golang.org/api/calendar/v3"
)

// type oauth2.Config struct {
// 	RedirectURL string

// }

type CalendarsApi struct {
	httpClient       http.Client
	calendarsService calendar.Service
}

func NewCalendarsApi(calendarService *calendar.Service) calendars.Repository {
	// http.HandleFunc("/GoogleCallback", handleGoogleCallback)
	return &CalendarsApi{
		httpClient:       http.Client{},
		calendarsService: calendar.Service{},
	}
}

func (ca *CalendarsApi) GetAll() (calendars.Event, error) {
	// calendarEvents, err := ca.calendarsService.Events.List("primary").TimeMin(time.Now().Format(time.RFC3339)).MaxResults(5).Do()
	// if err != nil {
	// 	// fmt.Fprintln(w, err)
	// 	return calendars.Event{}, err
	// }
	// if len(calendarEvents.Items) > 0 {
	// 	for _, i := range calendarEvents.Items {
	// 		fmt.Println(i.Summary, " ", i.Start.DateTime)
	// 	}
	// }
	fmt.Println("get all")
	return calendars.Event{}, nil
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
	if err != nil {
		return calendars.Event{}, err
	}

	eventDataUpdate := result
	// eventDataUpdate.Attendees
	requestLink := (fmt.Sprintf("https://www.googleapis.com/calendar/v3/calendars/%s/events", calendarId))
	reqBody, errUpdate := json.Marshal(eventDataUpdate)
	if errUpdate != nil {
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
