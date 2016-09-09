package routes

import (
	"context"
	"net/http"

	"github.com/tiborv/api-auth/db"
)

const cookieName = "_sess"

type contextKeys int

const (
	SessionCtxKey contextKeys = iota
	UserCtxKey
)

func SessionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, noCookie := r.Cookie(cookieName)
		if noCookie != nil {
			s := db.NewSession(cookieName)
			h.ServeHTTP(w, bindToRequest(w, r, s))
			return
		}
		s := db.FindSession(c.Value)
		h.ServeHTTP(w, bindToRequest(w, r, s))
	})
}

func bindToRequest(w http.ResponseWriter, r *http.Request, s db.Session) *http.Request {
	http.SetCookie(w, &http.Cookie{Name: cookieName, Value: s.Key, Expires: s.Expires})
	return r.WithContext(context.WithValue(r.Context(), SessionCtxKey, s))
}

func RequireUser(h http.HandlerFunc) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		//session := r.Context().Value(SessionCtxKey).(db.Session)
		h(w, r)

	})
}
