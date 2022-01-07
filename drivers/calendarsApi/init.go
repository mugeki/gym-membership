package calendarsApi

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

var (
	googleOauthConfig = &oauth2.Config{
		RedirectURL:  "http://localhost:8000/GoogleCallback",
		ClientID:     os.Getenv("CLIENT_ID"),    // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		ClientSecret: os.Getenv("googlesecret"), // from https://console.developers.google.com/project/<your-project-id>/apiui/credential
		Scopes:       []string{"https://www.googleapis.com/auth/calendar"},
		Endpoint:     google.Endpoint,
	}
	oauthStateString = "random"
	tokenString      = &oauth2.Token{}
)

func getToken(w http.ResponseWriter, r *http.Request) {
	state := r.FormValue("state")
	if state != oauthStateString {
		fmt.Printf("invalid oauth state, expected '%s', got '%s'\n", oauthStateString, state)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	code := r.FormValue("code")
	token, err := googleOauthConfig.Exchange(oauth2.NoContext, code)
	if err != nil {
		fmt.Printf("oauthConf.Exchange() failed with '%s'\n", err)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		// WillReturnRows
		return
	}
	tokenString = token
	// client := oauth2.NewClient(context.Background(), oauth2.StaticTokenSource(token))

}

func RequestToken() *oauth2.Token {
	http.HandleFunc("/GoogleLogin", getToken)
	fmt.Println("token =========== ", tokenString)
	tokenStashed, _ := json.Marshal(tokenString)
	fmt.Println(tokenStashed, "token stashed ======")
	return tokenString
}
