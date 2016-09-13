package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Token struct {
	Id       string   `json:"id"`
	Services []string `json:"services"`
}

const (
	tokenPrefix   = "TOKEN-"
	tokenIdLength = 30
)

func (t Token) Init() Token {
	t.Id = randomStringCrypto(tokenIdLength)
	return t
}

func NewToken() Token {
	return Token{}.Init()
}

func FindToken(id string) (Token, error) {
	jsonToken, err := redisClient.Get(id).Result()
	token := Token{}
	json.Unmarshal([]byte(jsonToken), &token)
	return token, err
}

func FindAllTokens() ([]Token, error) {
	tokens, err := redisClient.Keys(tokenPrefix + "*").Result()
	results := make([]Token, len(tokens))
	for i, t := range tokens {
		dbToken, _ := FindToken(t)
		results[i] = dbToken
	}
	return results, err
}

func (t Token) Save() (Token, error) {
	jsonToken, err := json.Marshal(t)
	if err != nil {
		fmt.Println("user serialization err:", err)
		return t, err
	}
	if t.Id == "" {
		fmt.Println("Token missing fields")
		return t, errors.New("Token missing fields")
	}
	redisClient.Set(tokenPrefix+t.Id, jsonToken, 0)
	return t, nil
}

func (t Token) AddService(serviceId string) (Token, error) {
	if stringInSlice(serviceId, t.Services) {
		fmt.Println("Service already associated with token")
		return t, errors.New("Service already associated with token")
	}
	t.Services = append(t.Services, serviceId)
	return t.Save()
}

func (t Token) Delete() bool {
	deleted, err := redisClient.Del(servicePrefix + t.Id).Result()
	if err != nil {
		fmt.Println("Service delete err: ", err)
		return false
	}
	return deleted > 0
}

func TokenJson(requestBody io.Reader) (Token, error) {
	token := Token{}
	err := json.NewDecoder(requestBody).Decode(&token)
	return token, err
}
