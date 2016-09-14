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
	mux.Handle(tokenPath+"/create", RequireUser(createToken))
	mux.Handle(tokenPath+"/delete", RequireUser(deleteToken))

}

func listToken(w http.ResponseWriter, r *http.Request) {
	tokens, _ := db.FindAllTokens()
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(tokens)
}

func createToken(w http.ResponseWriter, r *http.Request) {
	token, jsonErr := db.TokenJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Token create jsonErr:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not created")
		return
	}
	savedToken, saveErr := token.Init().Save()
	if saveErr != nil {
		fmt.Println("Token create err:", saveErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not created")
		return
	}
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(savedToken)
}

func updateToken(w http.ResponseWriter, r *http.Request) {
	token, jsonErr := db.TokenJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Token update jsonErr:", jsonErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not updated")
		return
	}
	savedToken, saveErr := token.Save()
	if saveErr != nil {
		fmt.Println("Token update err:", saveErr)
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(w, "Token not updated")
		return
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(savedToken)
}

func deleteToken(w http.ResponseWriter, r *http.Request) {
	token, jsonErr := db.TokenJson(r.Body)
	if jsonErr != nil {
		fmt.Println("Token delete jsonErr:", jsonErr)
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
