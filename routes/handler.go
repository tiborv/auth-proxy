package routes

import (
	"net/http"

	"github.com/urfave/negroni"

	db "../db"
)

func RootHandler() *negroni.Negroni {
	n := negroni.Classic()
	mux := http.NewServeMux()
	n.Use(negroni.HandlerFunc(db.SessionMiddleware))
	registerApiHandlers(mux)
	n.UseHandler(mux)

	return n

}
