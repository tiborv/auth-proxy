package db

import redis "gopkg.in/redis.v4"

var redisClient *redis.Client

func Connect() {
	redisClient = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

}
