package db

import (
	"encoding/json"
	"fmt"
)

type Service struct {
	Url string `json:"url"`
	Id  string `json:"id"`
}

const (
	servicePrefix    = "SERVICE-"
	serviceKeyLength = 8
)

func (s Service) InitService() Service {
	s.Id = generateRandomString(serviceKeyLength)
	return s
}

func NewService(url string) Service {
	return Service{Url: url}.InitService()
}

func (s Service) Save() Service {
	jsonService, err := json.Marshal(s)
	if err != nil {
		fmt.Println("user serialization err:", err)
		return s
	}

	redisClient.Set(servicePrefix+s.Id, jsonService, 0)
	return s
}

func FindService(id string) (Service, error) {
	jsonService, err := redisClient.Get(servicePrefix + id).Result()
	service := Service{}
	json.Unmarshal([]byte(jsonService), &service)
	return service, err
}

func CountServices() int {
	results, _ := redisClient.Keys(servicePrefix).Result()
	return len(results)
}

func (s Service) Exists() bool {
	exists, _ := redisClient.Exists(servicePrefix + s.Id).Result()
	return exists
}
