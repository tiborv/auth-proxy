package models

import (
	"os"

	redis "gopkg.in/redis.v4"
)

var redisClient *redis.Client

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

}
