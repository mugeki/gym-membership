package calendarsApi

import (
	"bytes"
	"encoding/json"
	"fmt"

	// "log"
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

// func getTokenAsync() <-chan *http.Request {
// 	token := make(chan *http.Request)
// 	go func() {
// 		defer close(token)
// 		time.Sleep(time.Second * 5)
// 		reqToken, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/google-login", nil)
// 		fmt.Println("++++++++", reqToken)
// 		// tokResult := calendars.Token{}
// 		// tokResult := "calendars.Token{}"
// 		token <- reqToken
// 	}()
// 	return token
// }

func (ca *CalendarsApi) GetAll(code, state string) (calendars.Event, error) {
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
	// reqToken, _ := http.NewRequest(http.MethodPost, "http://localhost:8000/google-login", nil)
	// fmt.Printf("===", reqToken)
	// var token interface{}
	// respToken, err := ca.httpClient.Do(reqToken)
	// fmt.Printf("===", respToken)
	// b := respToken.Body
	// fmt.Println("token string 1", b)
	// if err != nil {
	// 	fmt.Println("error get response token", err)
	// }
	// err = json.NewDecoder(respToken.Body).Decode(&token)
	// if err != nil {
	// 	fmt.Println(err)
	// }
	// if state != oauthStateString {
	// 	err := errors.New("invalid oauth state")
	// 	return calendars.Event{}, err
	// }
	// token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	// if err != nil {
	// 	fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
	// 	return calendars.Event{}, err
	// }

	// client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	// calendarService, err := calendar.New(client)

	// calendarEvents, err := calendarService.Events.List("primary").TimeMin(time.Now().Format(time.RFC3339)).MaxResults(5).Do()

	// // calendarService, err := calendar.New(token)
	// a := calendarEvents.Items
	// token := <-getTokenAsync()
	// fmt.Println("=====", tok)
	return calendars.Event{}, nil
}

func (ca *CalendarsApi) CreateEvent(eventData *calendars.Event) (calendars.Event, error) {
	calendarId := "elangmugeki@gmail.com" //tempCalendarID

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
	calendarId := "elangmugeki@gmail.com" //tempCalendarID

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
