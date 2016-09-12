package routes

import (
	"context"
	"net/http"

	"github.com/tiborv/prxy/db"
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
			s := db.NewSession().Save()
			h.ServeHTTP(w, bindToRequest(w, r, s))
			return
		}
		s, noSession := db.FindSession(c.Value)
		if noSession != nil {
			s := db.Session{}.InitSession(c.Value).Save()
			h.ServeHTTP(w, bindToRequest(w, r, s))
			return
		}
		h.ServeHTTP(w, bindToRequest(w, r, s))
	})
}

func bindToRequest(w http.ResponseWriter, r *http.Request, s db.Session) *http.Request {
	cookie := http.Cookie{Name: cookieName, Value: s.Key, Expires: s.Expires, Path: "/"}
	http.SetCookie(w, &cookie)
	return r.WithContext(context.WithValue(r.Context(), SessionCtxKey, s))
}

func GetSession(r *http.Request) db.Session {
	return r.Context().Value(SessionCtxKey).(db.Session)
}
