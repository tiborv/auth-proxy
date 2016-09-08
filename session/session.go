package session

import (
	"net/http"
	"time"
)

func GetOrSetSession(w http.ResponseWriter, r *http.Request) {
	expiration := time.Now().Add(365 * 24 * time.Hour)
	r.Cookies()
	cookie := http.Cookie{Name: "_sess", Value: "abcd", Expires: expiration}
	http.SetCookie(w, &cookie)

}
