package models

import (
	"fmt"
	"os"

	redis "gopkg.in/redis.v4"
)

var redisClient *redis.Client

const (
	adminUser     = "ADMIN"
	adminUserPass = "asdasdasd123"
)

func Connect() {

	redisUrl := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASS")
	if redisUrl == "" {
		redisUrl = "localhost:6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})

	_, noAdmin := FindUser(adminUser)

	if noAdmin != nil {
		fmt.Println("Create admin user")
		User{Username: adminUser, Password: []byte(adminUserPass)}.Save()
	}

}
