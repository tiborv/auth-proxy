package routes

import (
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/db"
)

const tokenPath = "/api/token"

func init() {
	mux.Handle(tokenPath+"/list", RequireUser(listTokens))
}

func listTokens(w http.ResponseWriter, r *http.Request) {
	tokens, err := db.FindAllTokens()
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Somthing went wrong")
		return
	}
	w.WriteHeader(http.StatusOK)
	fmt.Fprint(w, tokens)
}
