package routes

import (
	"fmt"
	"net/http"
	"regexp"

	"github.com/tiborv/api-auth/db"
)

func RequireUser(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(SessionCtxKey).(db.Session)
		if session.User != nil {
			h(w, r)
			return
		}
		w.WriteHeader(http.StatusUnauthorized)
		fmt.Fprint(w, "Not authenticated")

	})
}

var apiPrefix = regexp.MustCompile(`^\/api`)
var staicPrefix = regexp.MustCompile(`^\/static`)

func StaticFileMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case apiPrefix.MatchString(r.URL.Path):
			h.ServeHTTP(w, r)
		case staicPrefix.MatchString(r.URL.Path):
			http.ServeFile(w, r, r.URL.Path[1:])
		default:
			http.ServeFile(w, r, "./static/index.html")
		}
	})
}
