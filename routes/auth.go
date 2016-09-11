package routes

import (
	"fmt"
	"net/http"

	"github.com/tiborv/api-auth/db"
)

const authPath = "/api/auth"

func init() {
	mux.HandleFunc(authPath+"/login", login)
	mux.HandleFunc(authPath+"/logout", logout)
}

func login(w http.ResponseWriter, r *http.Request) {
	user, _ := db.UserJson(r.Body)
	s, err := GetSession(r).Auth(user)
	if err == nil {
		bindToRequest(w, r, s)
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Logged in")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Not logged in, bad username or password?")

}
func logout(w http.ResponseWriter, r *http.Request) {
	GetSession(r).RemoveUser().Save()
}
