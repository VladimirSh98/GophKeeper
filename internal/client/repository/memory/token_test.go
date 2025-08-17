package memory

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestTokenManager(t *testing.T) {
	secret := "mysecretkey1234567890123456"
	manager := NewManager(secret)

	home, err := os.UserHomeDir()
	require.NoError(t, err)

	dir := filepath.Join(home, ".gophkeeper")
	file := filepath.Join(dir, "token.enc")
	defer os.Remove(file)

	t.Run("Save and Get token", func(t *testing.T) {
		err := manager.SaveToken("mytoken123")
		require.NoError(t, err)

		token, err := manager.GetToken()
		require.NoError(t, err)
		require.Equal(t, "mytoken123", token)
	})

	t.Run("Expired token", func(t *testing.T) {
		data := map[string]interface{}{
			"token": "expiredtoken",
			"exp":   time.Now().Add(-1 * time.Hour).Format(time.RFC3339),
		}
		encrypted, err := encrypt(data, secret)
		require.NoError(t, err)
		require.NoError(t, os.WriteFile(file, encrypted, 0600))

		token, err := manager.GetToken()
		require.Error(t, err)
		require.Contains(t, err.Error(), "токен просрочен")
		require.Equal(t, "", token)
	})

	t.Run("Clear token", func(t *testing.T) {
		err := manager.SaveToken("tobedeleted")
		require.NoError(t, err)
		err = manager.ClearToken()
		require.NoError(t, err)
		_, err = os.Stat(file)
		require.True(t, os.IsNotExist(err))
	})
}

func TestEncrypt(t *testing.T) {
	type Sample struct {
		Field1 string
		Field2 int
	}

	data := Sample{
		Field1: "hello",
		Field2: 42,
	}
	secret := "mysecretkey1234567890123456"

	encrypted, err := encrypt(data, secret)
	require.NoError(t, err)
	require.NotEmpty(t, encrypted)

	key := make([]byte, 32)
	copy(key, []byte(secret))

	block, err := aes.NewCipher(key)
	require.NoError(t, err)

	gcm, err := cipher.NewGCM(block)
	require.NoError(t, err)

	nonceSize := gcm.NonceSize()
	require.GreaterOrEqual(t, len(encrypted), nonceSize)

	nonce, ciphertext := encrypted[:nonceSize], encrypted[nonceSize:]
	decrypted, err := gcm.Open(nil, nonce, ciphertext, nil)
	require.NoError(t, err)

	var result Sample
	err = json.Unmarshal(decrypted, &result)
	require.NoError(t, err)
	require.Equal(t, data, result)

	encrypted2, err := encrypt(data, secret)
	require.NoError(t, err)
	require.NotEqual(t, encrypted, encrypted2)
}

func TestDecrypt(t *testing.T) {
	secret := "mysecretkey1234567890123456"
	data := map[string]interface{}{
		"Field1": "hello",
		"Field2": float64(42),
	}

	ciphertext, err := encrypt(data, secret)
	require.NoError(t, err)
	require.NotEmpty(t, ciphertext)

	decrypted, err := decrypt(ciphertext, secret)
	require.NoError(t, err)
	require.Equal(t, data, decrypted)

	_, err = decrypt(ciphertext, "wrongsecretkeywrongsecretkey1234")
	require.Error(t, err)

	_, err = decrypt([]byte{0x01, 0x02}, secret)
	require.Error(t, err)
}
