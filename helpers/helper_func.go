package helpers

import (
	"math/rand"
	"strings"
	"time"
)

func RandomStrGeneration(needLength int) string {
	rand.Seed(time.Now().UnixNano())
	chars := []rune("ABCDEFGHIJKLMNOPQRSTUVWXYZ" +
		"abcdefghijklmnopqrstuvwxyz" +
		"0123456789")
	length := needLength
	var b strings.Builder
	for i := 0; i < length; i++ {
		b.WriteRune(chars[rand.Intn(len(chars))])
	}

	return strings.ToLower(b.String())
}
