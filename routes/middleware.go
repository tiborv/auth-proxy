package routes

import (
	"net/http"
	"strings"

	"github.com/tiborv/auth-proxy/models"
)

func RequireAuth(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		session := r.Context().Value(SessionCtxKey).(models.Session)
		if session.Auth {
			h(w, r)
			return
		}
		HttpResponse{Status: http.StatusUnauthorized, Msg: "Not authorized"}.Send(w)
	})
}

func StaticFileMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		switch {
		case strings.HasPrefix(r.URL.Path, "/static"):
			h.ServeHTTP(w, r)
			break
		case strings.HasPrefix(r.URL.Path, "/api"):
			w.Header().Set("Content-Type", "application/json")
			h.ServeHTTP(w, r)
			break
		default:
			http.ServeFile(w, r, "./static/index.html")
		}
	})
}
