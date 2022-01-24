package auth

import (
	"context"
	"errors"
	"fmt"
	controller "gym-membership/controllers"
	"gym-membership/controllers/auth/response"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/jinzhu/copier"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"

	// "gopkg.in/oauth2.v3/store"
	"golang.org/x/oauth2/google"
)

type AuthController struct {
	oauthConfig      *oauth2.Config
	oauthStateString string
	tokenString      *oauth2.Token
	calendarsService *calendar.Service
	client           *http.Client
}

func NewAuthController() *AuthController {
	return &AuthController{
		oauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:8000/google-callback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
			Endpoint:     google.Endpoint,
		},
		oauthStateString: "random",
		tokenString:      &oauth2.Token{},
		client:           &http.Client{},
		calendarsService: &calendar.Service{},
	}
}

func (ctrl *AuthController) HandleGoogleCallback(c echo.Context) error {
	state := c.FormValue("state")
	fmt.Println("state in callback function", state)
	if state != ctrl.oauthStateString {
		err := errors.New("invalid oauth state")
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	code := c.FormValue("code")
	fmt.Println("code in callback func", code)
	token, err := ctrl.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	calendarService, err := calendar.NewService(c.Request().Context(), option.WithHTTPClient(client))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	ctrl.client = client
	ctrl.tokenString = token
	ctrl.calendarsService = calendarService
	return controller.NewSuccessResponse(c, http.StatusOK, token)
}

func (ctrl *AuthController) HandleGoogleLogin(c echo.Context) error {
	authURL := ctrl.oauthConfig.AuthCodeURL(ctrl.oauthStateString)
	http.Redirect(c.Response(), c.Request(), authURL, http.StatusTemporaryRedirect)
	fmt.Println(authURL)
	return controller.NewSuccessResponse(c, http.StatusOK, authURL)
}

func (ctrl *AuthController) CreateEvent(calendarId, titleEvent, dateStart, dateEnd string) (response.Event, error) {
	calendarService := ctrl.calendarsService
	newEvent := calendar.Event{
		Summary: titleEvent,
		Start:   &calendar.EventDateTime{DateTime: dateStart, TimeZone: "Asia/Jakarta"},
		End:     &calendar.EventDateTime{DateTime: dateEnd, TimeZone: "Asia/Jakarta"},
	}
	createdEvent, err := calendarService.Events.Insert(calendarId, &newEvent).Do()
	if err != nil {
		fmt.Println("error create event", err)
		return response.Event{}, err
	}
	res := response.Event{}
	copier.Copy(&res, createdEvent)
	return res, nil
}

func (ctrl *AuthController) CreateCalendar(title string) (string, error) {
	calendarService := ctrl.calendarsService
	newCalendar := calendar.Calendar{
		Summary: title,
	}
	fmt.Println(newCalendar)
	createdCalendar, err := calendarService.Calendars.Insert(&newCalendar).Do()
	if err != nil {
		fmt.Println("error create calendar", err)
		return "", err
	}
	idCalendar := createdCalendar.Id
	return idCalendar, nil
}

func (ctrl *AuthController) CreatenewClassSchedule(c echo.Context) error {
	titleClass := "calisthenic 1 week program"
	listSchedule := "2022-01-25T07:00:00.000Z,2022-01-25T09:08:00.000Z;2022-01-26T07:00:00.000Z,2622-01-16T09:08:00.000Z;2022-01-27T07:00:00.000Z,2022-01-27T09:08:00.000Z"
	idCalendar, err := ctrl.CreateCalendar(titleClass)
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	listEventCreated := []response.Event{}
	eventRes := response.Event{}
	listDate := strings.Split(listSchedule, ";")
	for _, item := range listDate {
		rangeDate := strings.Split(item, ",")
		event, err := ctrl.CreateEvent(idCalendar, titleClass, rangeDate[0], rangeDate[1])
		if err != nil {
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
		copier.Copy(&eventRes, &event)
		listEventCreated = append(listEventCreated, eventRes)
	}

	return controller.NewSuccessResponse(c, http.StatusOK, listEventCreated)
}

func (ctrl *AuthController) UpdateAttendance(c echo.Context) error {
	calendarId := "ea053rm9emqclt6sb33t2irj4k@group.calendar.google.com" //temp calendar id
	email := "amirohqurrota98@gmail.com"                                 //temp email user
	calendarService := ctrl.calendarsService
	listEventUpdated := []response.Event{}
	events, err := calendarService.Events.List(calendarId).Do()
	for _, eventItem := range events.Items {
		idEvent := eventItem.Id
		attendees := eventItem.Attendees
		addAttendees := calendar.EventAttendee{
			Email: email,
		}
		reqAttend := append(attendees, &addAttendees)
		eventItem.Attendees = reqAttend
		updateEvent, err := calendarService.Events.Update(calendarId, idEvent, eventItem).Do()
		if err != nil {
			fmt.Println("error update event attendees", err)
			return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
		}
		resEvent := response.Event{}
		copier.Copy(&resEvent, &updateEvent)
		listEventUpdated = append(listEventUpdated, resEvent)
	}
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	return controller.NewSuccessResponse(c, http.StatusOK, events)
}
