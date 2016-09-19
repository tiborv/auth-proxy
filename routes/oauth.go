package routes

import (
	"context"
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/tiborv/auth-proxy/models"

	"golang.org/x/oauth2"
	"golang.org/x/oauth2/github"
)

type GitHubOrg struct {
	Login string
	Id    int
}

const (
	oauthPath       = "/api/oauth"
	gitHubOrgAPI    = "https://api.github.com/user/orgs"
	dbmedialabOrgID = 1803982
	soldontnoOrgID  = 501783
)

var config = &oauth2.Config{
	ClientID:     "d8b2829c1f6f3e67317a",
	ClientSecret: "d5ab38f3211bb36f21113bebc4e10dee5bb12439",
	Scopes:       []string{"read:org"},
	Endpoint:     github.Endpoint,
	RedirectURL:  "https://rocky-reef-55650.herokuapp.com/api/oauth/callback",
}

func init() {
	mux.HandleFunc(oauthPath+"/login", login)
	mux.HandleFunc(oauthPath+"/callback", callback)
	mux.HandleFunc(oauthPath+"/check", check)

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

	if tok.Valid() && UserInGitHubOrgs(tok, dbmedialabOrgID, soldontnoOrgID) {
		s := GetSession(r)
		s.Authenticate()
		bindSessionToCookie(w, r, s)
		http.Redirect(w, r, "/", http.StatusTemporaryRedirect)
		return
	}

	HttpResponse{Status: http.StatusBadRequest, Msg: "Could not authenticate"}.Send(w)
}

func UserInGitHubOrgs(tok *oauth2.Token, orgIDs ...int) bool {
	t := oauth2.Transport{Source: oauth2.StaticTokenSource(tok)}
	req, _ := http.NewRequest("GET", gitHubOrgAPI, nil)
	resp, _ := t.RoundTrip(req)
	userOrgs := []GitHubOrg{}
	json.NewDecoder(resp.Body).Decode(&userOrgs)
	for _, orgID := range orgIDs {
		for _, userOrg := range userOrgs {
			if userOrg.Id == orgID {
				return true
			}
		}
	}
	return false
}

func check(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(SessionCtxKey).(models.Session)
	HttpResponse{Status: http.StatusOK, Msg: strconv.FormatBool(session.Auth)}.Send(w)
}
