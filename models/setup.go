package models

import (
	"log"
	"os"

	mongo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v4"
)

var redisClient *redis.Client
var mongoClient *mongo.Database

const mongoDBName = "authProxy"

func Connect(mongoStats bool) {

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

	if mongoStats {
		mongoURL := os.Getenv("MONGO_URI")

		if mongoURL == "" {
			mongoURL = "localhost:27017"
		}
		session, err := mongo.Dial(mongoURL)
		if err != nil {
			log.Fatal(err)
		}
		session.SetMode(mongo.Monotonic, true)
		mongoClient = session.DB(mongoDBName)
	}

}
