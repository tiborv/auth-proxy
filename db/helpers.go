package db

import (
	crypto "crypto/rand"
	"encoding/base64"
	"math/rand"
)

func stringInSlice(a string, list []string) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func randomStringCrypto(n int) string {
	b := make([]byte, n)
	crypto.Read(b)
	return base64.URLEncoding.EncodeToString(b)
}

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890"

func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}
