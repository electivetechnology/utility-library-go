package hash

import (
	"math/rand"
)

const characters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func GenerateHash(length int) string {

	b := make([]byte, length)

	for i := range b {
		b[i] = characters[rand.Int63()%int64(len(characters))]
	}

	return string(b)
}
