package hash

import (
	"fmt"
	"testing"
	"unicode/utf8"
)

func TestGenerateHash(t *testing.T) {
	length := 10
	randomString := GenerateHash(length)

	if utf8.RuneCountInString(randomString) != length {
		t.Errorf("Length of hash does not equal to %d", length)
	}
}

func TestGenerateMultipleHashes(t *testing.T) {
	length := 10
	limit := 5

	hashArray := []string{}

	for i := 0; i < limit; i++ {
		randomString := GenerateHash(length)
		hashArray = append(hashArray, randomString)
		fmt.Println("Generated hash", randomString)
	}

	fmt.Println("hashArray", hashArray)

	for i := 0; i < limit; i++ {
		j := 0

		for ; j < i; j++ {
			if hashArray[j] == hashArray[i] {
				fmt.Println("hashArray[j]", hashArray[j])
				fmt.Println("hashArray[i]", hashArray[i])
				t.Errorf("%s\n is a duplicate %s\n", hashArray[j], hashArray[i])
			}
		}

	}
}
