package _utils

import (
	"math/rand"
	"strings"
	"time"
)

func RandomString(length int) string {
	rand.Seed(time.Now().Unix())
	var output strings.Builder
	charSet := "abcdedfghijklmnopqrstABCDEFGHIJKLMNOP"

	for i := 0; i < length; i++ {
		random := rand.Intn(len(charSet))
		randomChar := charSet[random]
		output.WriteString(string(randomChar))
	}
	return output.String()
}

func RandomInt(max int) int {
	return rand.Intn(max)
}
