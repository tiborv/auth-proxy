package routes

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/tiborv/prxy/db"
)

const tokenPath = "/api/token"

func init() {
	mux.Handle(tokenPath+"/list", RequireUser(listToken))
	mux.Handle(tokenPath+"/update", RequireUser(updateToken))
	mux.Handle(tokenPath+"/delete", RequireUser(deleteToken))
	mux.Handle(tokenPath+"/create", RequireUser(createToken))
}

func listToken(w http.ResponseWriter, r *http.Request) {
	tokens, _ := db.FindAllTokens()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tokens)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	token, saveErr := db.Token{}.Init().Save()
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(token)
}

func updateToken(w http.ResponseWriter, r *http.Request) {
	token, err := db.TokenJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not updated")
		return
	}
	_, saveErr := token.Save()
	if saveErr != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not updated")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(token)
}

func deleteToken(w http.ResponseWriter, r *http.Request) {
	token, err := db.TokenJson(r.Body)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not deleted")
		return
	}
	deleted := token.Delete()
	if deleted {
		w.WriteHeader(http.StatusOK)
		fmt.Fprint(w, "Token deleted")
		return
	}
	w.WriteHeader(http.StatusBadRequest)
	fmt.Fprint(w, "Token not deleted")

}
