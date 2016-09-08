package db

import (
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"golang.org/x/net/context"
)

type Session struct {
	Name    string
	Key     string
	Expires time.Time
	User    User
}

const sessionExpiration = 10 * time.Minute
const sessionPrefix = "SESSION-"
const cookieName = "_sess"

func SessionMiddleware(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {
	c, noCookie := r.Cookie(cookieName)
	if noCookie != nil {
		s := createSession()
		http.SetCookie(w, &http.Cookie{Name: s.Name, Value: s.Key, Expires: s.Expires})
		s.Save()
		next(w, r)
		return
	}
	s, noSession := findSession(c.Value)
	if noSession != nil {
		fmt.Println("no sess found")
		return
	}

	r = r.WithContext(context.WithValue(r.Context(), "user", s.User))
	next(w, r)
}

func (s Session) Save() (*Session, error) {
	sessionJSON, err := json.Marshal(s)
	if err != nil {
		return nil, err
	}
	redisClient.Set(sessionPrefix+s.Key, sessionJSON, 0)
	return &s, nil
}

func findSession(key string) (*Session, error) {
	sessionJSON, err := redisClient.Get(sessionPrefix + key).Result()
	if err != nil {
		return nil, err
	}
	session := Session{}
	json.Unmarshal([]byte(sessionJSON), &session)
	return &session, err
}

func createSession() Session {
	key := GenerateRandomString(30)
	expiration := time.Now().Add(sessionExpiration)
	session := Session{Name: cookieName, Key: key, Expires: expiration}
	return session
}

func GenerateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func GenerateRandomString(s int) string {
	b, _ := GenerateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)
}
