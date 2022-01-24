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
	"time"

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

func (ctrl *AuthController) GetAll(c echo.Context) error {
	calendarService := ctrl.calendarsService
	t := time.Now().Format(time.RFC3339)
	events, err := calendarService.Events.List("primary").ShowDeleted(false).
		SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
		log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}

	res := []response.Event{}
	copier.Copy(&res, events.Items)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *AuthController) CreateEvent(c echo.Context) error {
	calendarService := ctrl.calendarsService
	// start := time.Date(2022, 1, 24, 20, 24, 0, 0, time.Local).Format(time.RFC3339)
	// fmt.Println(start, "start time")
	newEvent := calendar.Event{
		Summary: "Testevent",
		Start:   &calendar.EventDateTime{DateTime: time.Date(2022, 1, 24, 20, 24, 0, 0, time.Local).Format(time.RFC3339)},
		End:     &calendar.EventDateTime{DateTime: time.Date(2022, 1, 24, 23, 21, 0, 0, time.Local).Format(time.RFC3339)},
	}
	createdEvent, err := calendarService.Events.Insert("primary", &newEvent).Do()
	if err != nil {
		fmt.Println("error create event", err)
		controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := response.Event{}
	copier.Copy(&res, createdEvent)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}

func (ctrl *AuthController) CreateCalendar(c echo.Context) error {
	idClass := 1 //temp id class
	calendarService := ctrl.calendarsService
	// start := time.Date(2022, 1, 24, 20, 24, 0, 0, time.Local).Format(time.RFC3339)
	// fmt.Println(start, "start time")
	newCalendar := calendar.Calendar{
		Id: string(idClass),
	}
	// newEvent := calendar.Event{
	// 	Summary: "Testevent",
	// 	Start:   &calendar.EventDateTime{DateTime: time.Date(2022, 1, 24, 20, 24, 0, 0, time.Local).Format(time.RFC3339)},
	// 	End:     &calendar.EventDateTime{DateTime: time.Date(2022, 1, 24, 23, 21, 0, 0, time.Local).Format(time.RFC3339)},
	// }
	// createdEvent, err := calendarService.Events.Insert("primary", &newEvent).Do()
	createdEvent, err := calendarService.Calendars.Insert(&newCalendar).Do()
	if err != nil {
		fmt.Println("error create event", err)
		controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	res := response.Event{}
	copier.Copy(&res, createdEvent)
	return controller.NewSuccessResponse(c, http.StatusOK, res)
}
