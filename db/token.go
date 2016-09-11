package db

import (
	"encoding/json"
	"fmt"
)

type Token struct {
	Id       string   `json:"id"`
	Services []string `json:"services"`
}

const (
	tokenPrefix   = "TOKEN-"
	tokenIdLength = 30
)

func (t Token) InitToken() Token {
	t.Id = generateRandomString(tokenIdLength)
	return t
}

func NewToken() Token {
	return Token{}.InitToken()
}

func FindToken(id string) (Token, error) {
	jsonToken, err := redisClient.Get(tokenPrefix + id).Result()
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

func (t Token) Save() Token {
	jsonToken, err := json.Marshal(t)
	if err != nil {
		fmt.Println("user serialization err:", err)
		return t
	}

	redisClient.Set(tokenPrefix+t.Id, jsonToken, 0)
	return t
}
