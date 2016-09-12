package db

import (
	"crypto/rand"
	"encoding/base64"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}

	return b, nil
}

func generateRandomString(n int) string {
	b, _ := generateRandomBytes(n)
	return base64.URLEncoding.EncodeToString(b)
}
