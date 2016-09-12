package routes

import (
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/db"
)

const servicePath = "/api/service"

func init() {
	mux.Handle(servicePath+"/create", RequireUser(createService))
}

func createService(w http.ResponseWriter, r *http.Request) {
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
