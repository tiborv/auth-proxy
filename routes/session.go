package routes

import (
	"context"
	"net/http"

	"github.com/tiborv/prxy/models"
)

type contextKeys int

const (
	cookieName                = "_sess"
	SessionCtxKey contextKeys = iota
	UserCtxKey
)

func SessionMiddleware(h http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		c, noCookie := r.Cookie(cookieName)
		if noCookie != nil {
			s := models.NewSession().Save()
			h.ServeHTTP(w, bindSessionToCookie(w, r, s))
			return
		}
		s, noSession := models.FindSession(c.Value)
		if noSession != nil {
			s := models.Session{}.InitSession(c.Value).Save()
			h.ServeHTTP(w, bindSessionToCookie(w, r, s))
			return
		}
		h.ServeHTTP(w, bindSessionToCookie(w, r, s))
	})
}

func bindSessionToCookie(w http.ResponseWriter, r *http.Request, s models.Session) *http.Request {
	cookie := http.Cookie{Name: cookieName, Value: s.Key, Expires: s.Expires, Path: "/"}
	http.SetCookie(w, &cookie)
	return r.WithContext(context.WithValue(r.Context(), SessionCtxKey, s))
}

func GetSession(r *http.Request) models.Session {
	return r.Context().Value(SessionCtxKey).(models.Session)
}
