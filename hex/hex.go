package hex

import (
	"crypto/rand"
	"encoding/hex"
)

func RandomHex(n int) (string, error) {
	bytes := make([]byte, n)
	if _, err := rand.Read(bytes); err != nil {
		return "", err
	}
	return hex.EncodeToString(bytes), nil
}

func IsValidHex(s string) bool {
	_, err := hex.DecodeString(s)
	return err == nil
}
