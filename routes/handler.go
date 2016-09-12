package routes

import "net/http"

var mux = http.NewServeMux()

func init() {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func GetRootMux() http.Handler {
	return SessionMiddleware(StaticFileMiddleware(mux))
}
