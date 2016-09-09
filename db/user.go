package db

import (
	"encoding/json"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string
	Password []byte
}

const userPrefix = "USER-"
const passwordCost = 8

func UserJson(requestBody io.Reader) (User, error) {
	user := User{}
	err := json.NewDecoder(requestBody).Decode(&user)
	return user, err
}

func (u User) Save() User {
	if u.AlreadyExists() || u.Password == nil {
		return u
	}
	u.Password, _ = bcrypt.GenerateFromPassword([]byte(u.Password), passwordCost)
	jsonUser, _ := json.Marshal(u)

	redisClient.Set(userPrefix+u.Username, jsonUser, 0)
	return u
}

func FindUser(username string) (User, error) {
	jsonUser, err := redisClient.Get(userPrefix + username).Result()
	user := User{}
	json.Unmarshal([]byte(jsonUser), &user)
	return user, err
}

func (u User) AlreadyExists() bool {
	_, err := redisClient.Get(userPrefix + u.Username).Result()
	return err == nil
}

func (u User) Delete() {
	redisClient.Del(userPrefix + u.Username)
	return
}

func (u *User) Auth() bool {
	dbUser, notFound := FindUser(u.Username)
	notCorrectPass := bcrypt.CompareHashAndPassword(dbUser.Password, u.Password)
	u.Password = dbUser.Password
	return notFound == nil && notCorrectPass == nil
}
