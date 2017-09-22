package utils

import (
	"math/rand"
	"time"
)

var seededRand *rand.Rand = rand.New(
	rand.NewSource(time.Now().UnixNano()))

// RandomString returns a random string with length = n
func RandomString(n int) string {
	var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ1234567890")

	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[seededRand.Intn(len(letterRunes))]
	}

	return string(b)
}
