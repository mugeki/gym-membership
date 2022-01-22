package auth

import (
	"context"
	"errors"
	controller "gym-membership/controllers"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/calendar/v3"
	"google.golang.org/api/option"
)

type AuthController struct{
	oauthConfig *oauth2.Config
	oauthStateString string
	tokenString *oauth2.Token
}

func NewAuthController() *AuthController {
	return &AuthController{
		oauthConfig: &oauth2.Config{
			RedirectURL:  "http://localhost:8000/GoogleCallback",
			ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),   
			ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
			Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
			Endpoint:     google.Endpoint,
		},
		oauthStateString: "random",
		tokenString     : &oauth2.Token{},
	}
}

// func (ctrl *CalendarsController) GetAll(c echo.Context) error {
// 	data, err := ctrl.calendarsUsecase.GetAll()
// 	if err != nil {
// 		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
// 	}
// 	return controller.NewSuccessResponse(c, http.StatusOK, data, nil)
// }

func (ctrl *AuthController) HandleGoogleLogin(c echo.Context) error {
	url := ctrl.oauthConfig.AuthCodeURL(ctrl.oauthStateString)
	http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	return controller.NewSuccessResponse(c, http.StatusOK, url)
}

func (ctrl *AuthController) HandleGoogleCallback(c echo.Context) error {
	state := c.FormValue("state")
	if state != ctrl.oauthStateString {
	
		// http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		err := errors.New("invalid oauth state")
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	code := c.FormValue("code")
	token, err := ctrl.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
	
		// http.Redirect(c.Response(), c.Request(), "/", http.StatusTemporaryRedirect)
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	
	calendarService, err := calendar.NewService(c.Request().Context(),option.WithHTTPClient(client))
	if err != nil {
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	// events, err := srv.Events.List("primary").ShowDeleted(false).
    //             SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
    //     if err != nil {
    //             log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
    //     }
	
	t := time.Now().Format(time.RFC3339)
	events, err := calendarService.Events.List("primary").ShowDeleted(false).
			SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	if err != nil {
			log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	}
	if len(events.Items) == 0 {
	} else {
			for _, item := range events.Items {
					date := item.Start.DateTime
					if date == "" {
							date = item.Start.Date
					}
			}
	}


	return c.JSON(http.StatusOK, client)
	// return controller.NewSuccessResponse(c, http.StatusOK, client)
}