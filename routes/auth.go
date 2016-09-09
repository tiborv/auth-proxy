package routes

import (
	"net/http"

	"github.com/tiborv/api-auth/db"
)

const authPath = "/auth"

func init() {
	mux.HandleFunc(authPath+"/login", login)
	mux.HandleFunc(authPath+"/logout", logout)

}

func login(w http.ResponseWriter, r *http.Request) {
	user, _ := db.UserJson(r.Body)
	session := r.Context().Value(SessionCtxKey).(db.Session)
	if user.Auth() {
		session.User = user
		session.Save()
	}

}
func logout(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(SessionCtxKey).(db.Session)
	session.User = db.User{}
	session.Save()
}
