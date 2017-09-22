package utils

import (
	"testing"
)

// TestRandomString checks that a created random strings are different and have a needed length
func TestRandomString(t *testing.T) {
	expectedLength := 20

	str1 := RandomString(expectedLength)
	str2 := RandomString(expectedLength)

	if len(str1) != expectedLength || len(str2) != expectedLength {
		t.Fatal("Random string does not have expected length")
	}

	if str1 == str2 {
		t.Fatal("Random strings are not random")
	}
}
