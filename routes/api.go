package routes

import (
	"fmt"
	"net/http"
)

func registerApiHandlers(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("/test", test)
	return mux
}

func test(w http.ResponseWriter, r *http.Request) {
	//user := db.User{Username: "swag", Password: "swag"}
	//user, err := db.FindUser("sawag")

	//c, _ := r.Cookie("_sess")
	fmt.Println(r.Context().Value("user"))

	//fmt.Print("test")
}
