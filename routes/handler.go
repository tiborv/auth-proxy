package routes

import (
	"encoding/json"
	"net/http"
)

var mux = http.NewServeMux()

func init() {
	mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))
}

func GetRootMux() http.Handler {
	return StaticFileMiddleware(SessionMiddleware(mux))
}

type HttpResponse struct {
	Msg    string `json:"msg"`
	Status int    `json:"status"`
}

func (he HttpResponse) Send(w http.ResponseWriter) {
	w.WriteHeader(he.Status)
	json.NewEncoder(w).Encode(he)
}
