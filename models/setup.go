package models

import (
	"log"
	"os"
	"time"

	mongo "gopkg.in/mgo.v2"
	redis "gopkg.in/redis.v4"
)

var redisClient *redis.Client

const mongoDBName = "auth-proxy"

var mongoStatsEnabled bool

func Connect(mongoStats bool) {
	mongoStatsEnabled = mongoStats

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

	if mongoStatsEnabled {
		mongoURL := os.Getenv("MONGODB_URI")

		if mongoURL == "" {
			mongoURL = "localhost:2701"
		}
		session, err := mongo.Dial(mongoURL)
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
