package memory

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/json"
	"errors"
	"io"
	"os"
	"path/filepath"
	"time"
)

// SaveToken save user token
func (t *Token) SaveToken(token string) error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	dir := filepath.Join(home, ".gophkeeper")
	if err = os.MkdirAll(dir, 0700); err != nil {
		return err
	}
	data := map[string]interface{}{
		"token": token,
		"exp":   time.Now().Add(24 * time.Hour).Format(time.RFC3339),
	}
	file := filepath.Join(dir, "token.enc")
	var encrypted []byte
	encrypted, err = encrypt(data, t.secretKey)
	if err != nil {
		return err
	}

	return os.WriteFile(file, encrypted, 0600)
}

// GetToken get user token
func (t *Token) GetToken() (string, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return "", err
	}
	file := filepath.Join(home, ".gophkeeper", "token.enc")
	var encrypted []byte
	encrypted, err = os.ReadFile(file)
	if err != nil {
		return "", err
	}
	var data map[string]interface{}
	data, err = decrypt(encrypted, t.secretKey)
	if err != nil {
		return "", err
	}
	expStr, ok := data["exp"].(string)
	if !ok {
		return "", errors.New("неверный формат времени в токене")
	}
	var expTime time.Time
	expTime, err = time.Parse(time.RFC3339, expStr)
	if err != nil {
		return "", errors.New("не удалось распарсить время токена")
	}
	if time.Now().After(expTime) {
		return "", errors.New("токен просрочен")
	}
	var token string
	token, ok = data["token"].(string)
	if !ok {
		return "", errors.New("неверный формат токена")
	}
	return token, nil
}

// ClearToken delete token
func (t *Token) ClearToken() error {
	home, err := os.UserHomeDir()
	if err != nil {
		return err
	}
	file := filepath.Join(home, ".gophkeeper", "token.enc")
	return os.Remove(file)
}

func encrypt(data interface{}, secret string) ([]byte, error) {
	key := make([]byte, 32)
	copy(key, []byte(secret))
	jsonData, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}
	var block cipher.Block
	block, err = aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var gcm cipher.AEAD
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	nonce := make([]byte, gcm.NonceSize())
	if _, err = io.ReadFull(rand.Reader, nonce); err != nil {
		return nil, err
	}
	return gcm.Seal(nonce, nonce, jsonData, nil), nil
}

func decrypt(ciphertext []byte, secret string) (map[string]interface{}, error) {
	key := make([]byte, 32)
	copy(key, []byte(secret))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	var gcm cipher.AEAD
	gcm, err = cipher.NewGCM(block)
	if err != nil {
		return nil, err
	}
	if len(ciphertext) < gcm.NonceSize() {
		return nil, errors.New("некорректный ciphertext")
	}
	nonce := ciphertext[:gcm.NonceSize()]
	data := ciphertext[gcm.NonceSize():]
	jsonData, err := gcm.Open(nil, nonce, data, nil)
	if err != nil {
		return nil, err
	}
	var result map[string]interface{}
	if err = json.Unmarshal(jsonData, &result); err != nil {
		return nil, err
	}
	return result, nil
}
