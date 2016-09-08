package db

import "encoding/json"

type User struct {
	Username string
	Password string
}

const userPrefix = "USER-"

func (u User) Save() User {
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
