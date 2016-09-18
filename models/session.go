package models

import (
	"encoding/json"
	"log"
	"time"
)

type Session struct {
	Key     string
	Expires time.Time
	Auth    bool
}

const (
	sessionDuration  = 10 * time.Minute
	sessionPrefix    = "SESSION-"
	sessionKeyLength = 64
)

func NewSession() Session {
	return Session{}.InitSession(randomStringCrypto(sessionKeyLength))
}

func (s Session) InitSession(key string) Session {
	s.Key = key
	s.Expires = time.Now().Add(sessionDuration)
	return s
}

func FindSession(key string) (Session, error) {
	sessionJSON, err := redisClient.Get(sessionPrefix + key).Result()
	if err != nil {
		return Session{}, err
	}
	session := Session{}
	json.Unmarshal([]byte(sessionJSON), &session)
	return session, nil
}

func (s Session) Save() Session {
	sessionJSON, jsonErr := json.Marshal(s)
	if jsonErr != nil {
		log.Println(jsonErr)
	}
	redisClient.Set(sessionPrefix+s.Key, sessionJSON, s.Expires.Sub(time.Now()))
	return s
}

func (s Session) Authenticate() Session {
	s.Auth = true
	return s.Save()
}
