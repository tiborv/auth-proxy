package models

import (
	"fmt"
	"log"
	"os"
	"time"

	mongo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v5"
)

var redisClient *redis.Client

const mongoDBName = "auth-proxy"

var mongoStatsEnabled bool

func Connect(mongoStats bool) {
	mongoStatsEnabled = mongoStats

	redisUrl := os.Getenv("REDIS_URL")
	redisPass := os.Getenv("REDIS_PASS")
	if redisUrl == "" {
		fmt.Println("Connecting to localhost redis")
		redisUrl = "localhost:6379"
	}

	redisClient = redis.NewClient(&redis.Options{
		Addr:     redisUrl,
		Password: redisPass, // no password set
		DB:       0,         // use default DB
	})

	if mongoStatsEnabled {
		mongoDBURI := os.Getenv("MONGO_URI")

		if mongoDBURI == "" {
			mongoDBURI = "localhost:27017"
		}

		log.Println("Connecting mongodb: ", mongoDBURI)
		session, err := mongo.Dial(mongoDBURI)
		if err != nil {
			log.Fatal("MONGO ERR", err)
		}
		session.SetMode(mongo.Monotonic, true)
		db := session.DB(mongoDBName)
		requestCollection = db.C("request")
		responseCollection = db.C("response")

		expirationIndex := mongo.Index{
			ExpireAfter: time.Duration(60*60) * time.Second,
			Unique:      true,
			DropDups:    true,
			Background:  true,
			Sparse:      true,
		}

		requestCollection.EnsureIndex(expirationIndex)
		responseCollection.EnsureIndex(expirationIndex)
	}
}
