package utils

import (
	"math/rand"
	"time"
)

const (
	ShortCodeLength = 7
	charset         = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"
)

func GenerateShortCode(lenth int) string {
	rand.New(rand.NewSource(time.Now().UnixNano()))
	shortCode := make([]byte, lenth)
	for i := range shortCode {
		shortCode[i] = charset[rand.Intn(len(charset))]
	}

	return string(shortCode)
}
