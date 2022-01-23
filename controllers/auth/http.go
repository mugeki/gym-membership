package auth

import (
	"context"
	"errors"
	"fmt"
	controller "gym-membership/controllers"
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"

	// "gopkg.in/oauth2.v3/store"
	"golang.org/x/oauth2/google"
)

type AuthController struct {
	oauthConfig      *oauth2.Config
	oauthStateString string
	tokenString      *oauth2.Token
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
	// url := ctrl.oauthConfig.AuthCodeURL(ctrl.oauthStateString)
	// fmt.Println(url, "url login")
	// http.Redirect(c.Response(), c.Request(), url, http.StatusTemporaryRedirect)
	// return controller.NewSuccessResponse(c, http.StatusOK, url)
	authURL := ctrl.oauthConfig.AuthCodeURL("state-token", oauth2.AccessTypeOffline)
	fmt.Printf("Go to the following link in your browser then type the "+
		"authorization code: \n%v\n", authURL)

	var authCode string
	if _, err := fmt.Scan(&authCode); err != nil {
		log.Fatalf("Unable to read authorization code: %v", err)
	}

	tok, _ := ctrl.oauthConfig.Exchange(context.TODO(), authCode)
	fmt.Println("token stirng", tok)
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve token from web: %v", err)
	// }
	return controller.NewSuccessResponse(c, http.StatusOK, tok)
}

func (ctrl *AuthController) HandleGoogleCallback(c echo.Context) error {
	state := c.FormValue("state")
	if state != ctrl.oauthStateString {
		err := errors.New("invalid oauth state")
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}

	code := c.FormValue("code")
	token, err := ctrl.oauthConfig.Exchange(context.Background(), code)
	if err != nil {
		// fmt.Println("error get token")
		return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	}
	// client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
	// calendarService, err := calendar.NewService(c.Request().Context(), option.WithHTTPClient(client))
	// if err != nil {
	// 	fmt.Println("error get service")
	// 	return controller.NewErrorResponse(c, http.StatusInternalServerError, err)
	// }

	// t := time.Now().Format(time.RFC3339)
	// events, err := calendarService.Events.List("primary").ShowDeleted(false).
	// 	SingleEvents(true).TimeMin(t).MaxResults(10).OrderBy("startTime").Do()
	// if err != nil {
	// 	log.Fatalf("Unable to retrieve next ten of the user's events: %v", err)
	// }
	// if len(events.Items) == 0 {
	// } else {
	// 	for _, item := range events.Items {
	// 		fmt.Println(item, "itemm")
	// 		date := item.Start.DateTime
	// 		if date == "" {
	// 			date = item.Start.Date
	// 		}
	// 	}
	// }
	return controller.NewSuccessResponse(c, http.StatusOK, token)
	// return calendarService, err
}
