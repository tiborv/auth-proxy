package models

import (
	"encoding/json"
	"errors"
	"io"
)

type Service struct {
	Url     string   `json:"url"`
	Host    string   `json:"host"`
	Scheme  string   `json:"scheme"`
	Slug    string   `json:"slug"`
	Clients []string `json:"clients"`
}

const (
	servicePrefix    = "SERVICE-"
	tokenPrefix      = "TOKENS-"
	serviceKeyLength = 16
)

func NewService(url, slug string) Service {
	return Service{Url: url, Slug: slug}
}

func (s Service) Save() (Service, error) {
	jsonService, err := json.Marshal(s)
	if err != nil {
		return s, err
	}
	if s.Slug == "" || s.Host == "" || s.Url == "" {
		return s, errors.New("Service missing fields")
	}
	if s.Scheme == "" {
		s.Scheme = "http"
	}
	redisClient.Del(tokenPrefix + s.Slug).Result()
	for _, c := range s.Clients {
		redisClient.SAdd(tokenPrefix+s.Slug, c)
	}
	redisClient.Set(servicePrefix+s.Slug, jsonService, 0)
	return s, nil
}

func FindService(redisKey string) (Service, error) {
	return FindServiceBySlug(redisKey[len(servicePrefix):])
}

func FindServiceBySlug(slug string) (Service, error) {
	jsonService, err := redisClient.Get(servicePrefix + slug).Result()
	service := Service{}
	json.Unmarshal([]byte(jsonService), &service)
	clients, _ := redisClient.SMembers(tokenPrefix + slug).Result()
	service.Clients = clients
	return service, err
}

func (s Service) Exists() bool {
	exists, _ := redisClient.Exists(servicePrefix + s.Slug).Result()
	return exists
}

func ServiceJson(requestBody io.Reader) (Service, error) {
	service := Service{}
	err := json.NewDecoder(requestBody).Decode(&service)
	return service, err
}

func (s Service) Delete() bool {
	servicesDeleted, err := redisClient.Del(servicePrefix + s.Slug).Result()
	_, errSet := redisClient.Del(tokenPrefix + s.Slug).Result()

	if err != nil || errSet != nil {
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

func GetUrlOfService(slug string) (string, error) {
	service, err := FindServiceBySlug(slug)
	return service.Url, err
}

func ServiceHasToken(serviceSlug, token string) (bool, error) {
	return redisClient.SIsMember(tokenPrefix+serviceSlug, token).Result()
}

func (s Service) RemoveToken(token string) (bool, error) {
	removed, err := redisClient.SRem(tokenPrefix+s.Slug, token).Result()
	return removed > 0, err
}
