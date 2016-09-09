package routes

import (
	"fmt"
	"net/http"

	"github.com/tiborv/api-auth/db"
)

const apiPath = "/api"

func init() {
	mux.Handle(apiPath+"/user/create", RequireUser(createUser))
	mux.Handle(apiPath+"/user/show", RequireUser(printUser))
}

func createUser(w http.ResponseWriter, r *http.Request) {
	user, err := db.UserJson(r.Body)
	if err != nil {
		fmt.Println(w, "Could not create user")
		return
	}
	user.Save()
}

func printUser(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(SessionCtxKey).(db.Session)
	fmt.Println(session.User)

}
