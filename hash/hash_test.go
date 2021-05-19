package hash

import (
	"testing"
	"unicode/utf8"
)

func TestGenerateHash(t *testing.T) {
	length := 10
	generatedString := GenerateHash(length)

	if utf8.RuneCountInString(generatedString) != length {
		t.Errorf("Length of hash does not equal to %d", length)
	}
}

func TestGenerateMultipleHashes(t *testing.T) {
	length := 10
	limit := 5

	hashArray := []string{}

	for i := 0; i < limit; i++ {
		generatedString := GenerateHash(length)
		hashArray = append(hashArray, generatedString)
	}

	for i := 0; i < limit; i++ {
		j := 0

		for ; j < i; j++ {
			if hashArray[j] == hashArray[i] {
				t.Errorf("%s\n is a duplicate of %s\n", hashArray[j], hashArray[i])
			}
		}

	}
}
