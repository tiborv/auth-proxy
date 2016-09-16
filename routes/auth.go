package routes

import (
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/models"
)

const authPath = "/api/auth"

func init() {
	mux.HandleFunc(authPath+"/login", login)
	mux.HandleFunc(authPath+"/logout", logout)
}

func login(w http.ResponseWriter, r *http.Request) {
	user, _ := models.UserJson(r.Body)
	s, err := GetSession(r).Auth(user)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Println(err)
		fmt.Fprint(w, "Not logged in, bad username or password?")
		return
	}
	bindToRequest(w, r, s)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "Logged in")
}

func logout(w http.ResponseWriter, r *http.Request) {
	GetSession(r).RemoveUser().Save()
}
