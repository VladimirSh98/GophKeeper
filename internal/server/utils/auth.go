package utils

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/base64"
	"fmt"
	"strings"
)

type contextKey string

const UserLoginKey contextKey = "Login"

const (
	saltSize       = 16
	hashIterations = 10
)

// GenerateSalt создает соль
func GenerateSalt() (string, error) {
	salt := make([]byte, saltSize)
	_, err := rand.Read(salt)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(salt), nil
}

// HashPassword создает комбинированную строку: "соль:хеш"
func HashPassword(password string) (string, error) {
	salt, err := GenerateSalt()
	if err != nil {
		return "", fmt.Errorf("failed to generate salt: %v", err)
	}
	hash := sha256.Sum256([]byte(password + salt))
	for i := 0; i < hashIterations-1; i++ {
		hash = sha256.Sum256(hash[:])
	}
	return fmt.Sprintf("%s:%x", salt, hash), nil
}

// VerifyPassword проверяет пароль
func VerifyPassword(password string, combinedHash string) bool {
	parts := strings.Split(combinedHash, ":")
	if len(parts) != 2 {
		return false
	}
	salt := parts[0]
	storedHash := parts[1]

	hash := sha256.Sum256([]byte(password + salt))
	for i := 0; i < hashIterations-1; i++ {
		hash = sha256.Sum256(hash[:])
	}

	return fmt.Sprintf("%x", hash) == storedHash
}
