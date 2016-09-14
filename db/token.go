package db

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Token struct {
	Name     string   `json:"name"`
	Id       string   `json:"id"`
	Services []string `json:"services"`
}

const (
	tokenPrefix   = "TOKEN-"
	tokenIdLength = 64
)

func (t Token) Init() Token {
	t.Id = randomStringCrypto(tokenIdLength)
	return t
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
	for i, tid := range tokens {
		dbToken, _ := FindToken(tid)
		results[i] = dbToken
	}
	return results, err
}

func (t Token) Save() (Token, error) {
	jsonToken, err := json.Marshal(t)
	if err != nil {
		fmt.Println("Token serialization err:", err)
		return t, err
	}
	if t.Id == "" || t.Name == "" {
		fmt.Println("Token missing fields")
		return t, errors.New("Token missing fields")
	}
	redisClient.Set(tokenPrefix+t.Id, jsonToken, 0)
	return t, nil
}

func (t Token) Delete() bool {
	deleted, err := redisClient.Del(tokenPrefix + t.Id).Result()
	if err != nil {
		fmt.Println("Token delete err: ", err)
		return false
	}
	return deleted > 0
}

func TokenJson(requestBody io.Reader) (Token, error) {
	token := Token{}
	err := json.NewDecoder(requestBody).Decode(&token)
	return token, err
}
