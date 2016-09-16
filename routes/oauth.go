package routes

import (
	"context"
	"log"
	"net/http"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

const oauthPath = "/api/oauth"

var config = &oauth2.Config{
	ClientID:     "d8b2829c1f6f3e67317a",
	ClientSecret: "d5ab38f3211bb36f21113bebc4e10dee5bb12439",
	Scopes:       []string{"read:org"},
	Endpoint:     github.Endpoint,
	RedirectURL:  "http://localhost:3000/api/oauth/callback",
}

func init() {
	mux.HandleFunc(oauthPath+"/login", login)
	mux.HandleFunc(oauthPath+"/callback", callback)

}

func login(w http.ResponseWriter, r *http.Request) {
	url := config.AuthCodeURL("state", oauth2.AccessTypeOffline)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func callback(w http.ResponseWriter, r *http.Request) {
	ctx := context.Background()
	tok, err := config.Exchange(ctx, r.URL.Query().Get("code"))
	if err != nil {
		log.Fatal(err)
	}

	if tok.Valid() {
		s := GetSession(r)
		s.Authenticate()
		bindToRequest(w, r, s)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)

	}

}
