package db

import (
	"encoding/json"
	"errors"
	"io"
)

type Service struct {
	Url    string   `json:"url"`
	Id     string   `json:"id"`
	Name   string   `json:"name"`
	Tokens []string `json:"tokens"`
}

const (
	servicePrefix    = "SERVICE-"
	tokenPrefix      = "TOKENS-"
	serviceKeyLength = 16
)

func (s Service) Init() Service {
	s.Id = randomString(serviceKeyLength)
	return s
}

func NewService(url, name string) Service {
	return Service{Url: url, Name: name}.Init()
}

func (s Service) Save() (Service, error) {
	jsonService, err := json.Marshal(s)
	if err != nil {
		return s, err
	}
	if s.Id == "" || s.Name == "" || s.Url == "" {
		return s, errors.New("Service missing fields")
	}
	redisClient.Del(tokenPrefix + s.Id)
	for _, t := range s.Tokens {
		redisClient.SAdd(tokenPrefix+s.Id, t)
	}
	redisClient.Set(servicePrefix+s.Id, jsonService, 0)
	return s, nil
}

func FindService(redisId string) (Service, error) {
	return FindServiceById(redisId[len(servicePrefix):])
}

func FindServiceById(id string) (Service, error) {
	jsonService, err := redisClient.Get(servicePrefix + id).Result()
	service := Service{}
	json.Unmarshal([]byte(jsonService), &service)

	tokens, _ := redisClient.SMembers(tokenPrefix + id).Result()
	service.Tokens = tokens
	return service, err
}

func (s Service) Exists() bool {
	exists, _ := redisClient.Exists(servicePrefix + s.Id).Result()
	return exists
}

func ServiceJson(requestBody io.Reader) (Service, error) {
	service := Service{}
	err := json.NewDecoder(requestBody).Decode(&service)
	return service, err
}

func (s Service) Delete() bool {
	servicesDeleted, err := redisClient.Del(servicePrefix + s.Id).Result()
	if err != nil {
		return false
	}
	return servicesDeleted > 0
}

func FindAllServices() ([]Service, error) {
	services, err := redisClient.Keys(servicePrefix + "*").Result()
	if err != nil {
		return []Service{}, err
	}
	return PopulateServices(services)
}

func PopulateServices(services []string) ([]Service, error) {
	result := make([]Service, len(services))
	for i, s := range services {
		dbService, err := FindService(s)
		if err != nil {
			return []Service{}, err
		}
		result[i] = dbService
	}
	return result, nil
}

func GetUrlOfService(id string) (string, error) {
	service, err := FindServiceById(id)
	return service.Url, err
}
