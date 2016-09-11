package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/api-auth/db"
)

const (
	userPath = "/api/user"
)

func init() {
	mux.Handle(userPath+"/create", RequireUser(createUser))
	mux.Handle(userPath+"/show", RequireUser(showUser))
	mux.Handle(userPath+"/delete", RequireUser(deleteUser))
	mux.Handle(userPath+"/list", RequireUser(listUsers))

}

func createUser(w http.ResponseWriter, r *http.Request) {
	user, err := db.UserJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User not created")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User created")
	user.Save()
}

func showUser(w http.ResponseWriter, r *http.Request) {
	session := r.Context().Value(SessionCtxKey).(db.Session)
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, session.User)
}

func deleteUser(w http.ResponseWriter, r *http.Request) {
	user, err := db.UserJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "User not deleted")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, "User deleted")
	user.Delete()
}

func listUsers(w http.ResponseWriter, r *http.Request) {
	users, err := db.FindAllUsers()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Somthing went wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(users)
}
