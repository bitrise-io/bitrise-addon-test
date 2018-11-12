package utils

import (
	"encoding/hex"
	"math/rand"
	"time"
)

// RandomHex ...
func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	rand.Seed(time.Now().UnixNano())
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}
