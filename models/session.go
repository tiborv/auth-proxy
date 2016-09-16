package models

import (
	"encoding/json"
	"fmt"
	"log"
	"time"

	"golang.org/x/crypto/bcrypt"
)

type Session struct {
	Key     string
	Expires time.Time
	User    *User
}

const (
	sessionDuration  = 30 * time.Minute
	sessionPrefix    = "SESSION-"
	sessionKeyLength = 30
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

func (s Session) Auth(u User) (Session, error) {
	dbUser, notFound := FindUser(u.Username)
	if notFound != nil {
		fmt.Println("Auth: user not found")
		return s, notFound
	}
	notCorrectPass := bcrypt.CompareHashAndPassword(dbUser.Password, u.Password)
	if notFound != nil {
		fmt.Println("Auth: user not correct pass")
		return s, notCorrectPass
	}
	s.User = &dbUser
	return s.Save(), nil

}

func (s Session) SetUser(u User) Session {
	s.User = &u
	return s.Save()
}

func (s Session) RemoveUser() Session {
	s.User = nil
	return s.Save()
}
