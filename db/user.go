package db

import (
	"encoding/json"
	"fmt"
	"io"

	"golang.org/x/crypto/bcrypt"
)

type User struct {
	Username string `json:"username"`
	Password []byte `json:"password"`
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
		fmt.Println("User not saved")
		return u
	}
	u.Password, _ = bcrypt.GenerateFromPassword([]byte(u.Password), passwordCost)
	jsonUser, err := json.Marshal(u)
	if err != nil {
		fmt.Println("user serialization err:", err)
		return u
	}

	redisClient.Set(userPrefix+u.Username, jsonUser, 0)
	return u
}

func FindUser(username string) (User, error) {
	jsonUser, err := redisClient.Get(userPrefix + username).Result()
	user := User{}
	json.Unmarshal([]byte(jsonUser), &user)
	return user, err
}

func FindAllUsers() ([]User, error) {
	users, err := redisClient.Keys(userPrefix + "*").Result()
	result := make([]User, len(users))
	for i, u := range users {
		dbUser, _ := FindUser(u)
		result[i] = dbUser
	}
	return result, err
}

func (u User) AlreadyExists() bool {
	exists, _ := redisClient.Exists(userPrefix + u.Username).Result()
	return exists
}

func (u User) Delete() {
	redisClient.Del(userPrefix + u.Username)
	return
}
