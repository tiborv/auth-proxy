package db

import (
	"encoding/base64"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

type Session struct {
	Name    string
	Key     string
	Expires time.Time
	User    User
}

const sessionDuration = 30 * time.Minute

const sessionPrefix = "SESSION-"
const keyLength = 30

func (s Session) Save() Session {
	sessionJSON, err := json.Marshal(s)
	if err != nil {
		log.Println(err)
	}
	redisClient.Set(sessionPrefix+s.Key, sessionJSON, 0)
	return s
}

func FindSession(key string) Session {
	sessionJSON, err := redisClient.Get(sessionPrefix + key).Result()
	if err != nil {
		log.Println(err)
	}
	session := Session{}
	json.Unmarshal([]byte(sessionJSON), &session)
	return session
}

func NewSession(name string) Session {
	return Session{Name: name, Key: generateRandomString(keyLength), Expires: time.Now().Add(sessionDuration)}.Save()
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	// Note that err == nil only if we read len(b) bytes.
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(s int) string {
	b, _ := generateRandomBytes(s)
	return base64.URLEncoding.EncodeToString(b)
}
