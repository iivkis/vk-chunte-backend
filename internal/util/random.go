package util

import (
	"math/rand"
	"time"
)

func init() {
	rand.Seed(time.Now().UnixNano())
}

const alphabet = "asdfghjklzxcvbnmqwertyuiop"

func GetRandomString(size int) string {
	var b []byte = make([]byte, size)
	for i := range b {
		b[i] = alphabet[rand.Intn(len(alphabet))]
	}
	return string(b)
}

func GetRandomInt(min int, max int) int {
	return min + rand.Intn(max-min+1)
}
