package utils

import (
	"math/rand"
	"strings"
	"time"
)

const (
	alphabet = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789_"
	length   = 10
)

func GenerateShortURL() string {
	rand.NewSource(time.Now().UnixNano())

	var sb strings.Builder
	for i := 0; i < length; i++ {
		randomIndex := rand.Intn(len(alphabet))
		sb.WriteByte(alphabet[randomIndex])
	}

	return sb.String()
}
