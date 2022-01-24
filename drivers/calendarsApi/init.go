package calendarsApi

import (
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/google-callback",
		ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
		ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET"),
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
	tokenString      = &oauth2.Token{}
)

// calendarService, err := calendar.New(client)

// func getToken(w http.ResponseWriter, r *http.Request) *oauth2.Token {
// 	state := r.FormValue("state")
// 	if state != oauthStateString {
// 		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		return nil
// 	}

// 	code := r.FormValue("code")
// 	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
// 	if err != nil {
// 		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
// 		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
// 		// WillReturnRows
// 		return nil
// 	}
// 	tokenString = token
// 	// client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))
// 	return tokenString
// 	// client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

// }

// // func RequestToken() *oauth2.Token {
// // 	url := ctrl.oauthConfig.AuthCodeURL(ctrl.oauthStateString)
// // 	http.HandleFunc("/GoogleLogin", getToken)
// // 	fmt.Println("token =========== ", tokenString)
// // 	tokenStashed, _ := json.Marshal(tokenString)
// // 	fmt.Println(tokenStashed, "token stashed ======")
// // 	return tokenString
// // }
