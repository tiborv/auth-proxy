package db

import (
	"fmt"

	redis "gopkg.in/redis.v4"
)

var redisClient *redis.Client

const adminUser = "ADMIN"
const adminUserPass = "asdasdasd123"

func Connect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	_, noAdmin := FindUser(adminUser)

	if noAdmin != nil {
		fmt.Println("Create admin user")
		User{Username: adminUser, Password: []byte(adminUserPass)}.Save()
	}

}
