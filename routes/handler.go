package routes

import "net/http"

var mux = http.NewServeMux()

func GetRootMux() http.Handler {
	return SessionMiddleware(mux)
}
