package models

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
)

type Client struct {
	Name  string       `json:"name"`
	Token string       `json:"token"`
	Stats RequestStats `json:"stats"`
}

const (
	clientPrefix   = "Client-"
	clientIdLength = 64
)

func (c Client) Init() Client {
	c.Token = randomStringCrypto(clientIdLength)
	return c
}

func FindClient(id string) (Client, error) {
	jsonClient, err := redisClient.Get(id).Result()
	client := Client{}
	json.Unmarshal([]byte(jsonClient), &client)
	client.GetRequestStats()
	return client, err
}

func FindAllClients() ([]Client, error) {
	clients, err := redisClient.Keys(clientPrefix + "*").Result()
	results := make([]Client, len(clients))
	for i, tid := range clients {
		dbClient, _ := FindClient(tid)
		results[i] = dbClient
	}
	return results, err
}

func (c Client) Save() (Client, error) {
	jsonClient, err := json.Marshal(c)
	if err != nil {
		fmt.Println("Client serialization err:", err)
		return c, err
	}
	if c.Token == "" || c.Name == "" {
		fmt.Println("Client missing fields")
		return c, errors.New("Client missing fields")
	}
	redisClient.Set(clientPrefix+c.Token, jsonClient, 0)
	return c, nil
}

func (c Client) Delete() bool {
	deleted, err := redisClient.Del(clientPrefix + c.Token).Result()
	if err != nil {
		fmt.Println("Client delete err: ", err)
		return false
	}
	return deleted > 0
}

func ClientJson(requestBody io.Reader) (Client, error) {
	Client := Client{}
	err := json.NewDecoder(requestBody).Decode(&Client)
	return Client, err
}
