package hash

import (
	"math/rand"
	"time"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateHash(length int) string {
	// Add random seed
	rand.Seed(time.Now().UnixNano())

	b := make([]byte, length)

	for i := range b {
		b[i] = characters[rand.Int63()%int64(len(characters))]
	}

	return string(b)
}
