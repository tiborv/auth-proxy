package routes

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/tiborv/prxy/db"
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

func StaticFileMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, "/static") || strings.HasPrefix(r.URL.Path, "/api") {
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			return
		}
		http.ServeFile(w, r, "./static/index.html")

	})
}
