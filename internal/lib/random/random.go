package random

import (
	"math/rand"
	"strings"
	"time"
)

const symbols = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func NewRandomString(length int) string {
	source := rand.NewSource(time.Now().UnixNano())
	r := rand.New(source)

	var sb strings.Builder
	sb.Grow(length)

	for range length {
		randomIndex := r.Intn(len(symbols))
		randomChar := symbols[randomIndex]
		sb.WriteByte(randomChar)
	}

	return sb.String()
}