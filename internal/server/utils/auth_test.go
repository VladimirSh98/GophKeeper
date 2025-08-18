package utils

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestGenerateSalt(t *testing.T) {
	salt, err := GenerateSalt()
	require.NoError(t, err)
	require.NotEmpty(t, salt)
	require.GreaterOrEqual(t, len(salt), 16)
}

func TestHashPasswordAndVerifyPassword(t *testing.T) {
	password := "mySecret123!"

	combinedHash, err := HashPassword(password)
	require.NoError(t, err)
	require.Contains(t, combinedHash, ":")

	valid := VerifyPassword(password, combinedHash)
	require.True(t, valid)

	invalid := VerifyPassword("wrongPassword", combinedHash)
	require.False(t, invalid)

	broken := VerifyPassword(password, "invalidformat")
	require.False(t, broken)
}

func TestHashPassword_UniqueSalts(t *testing.T) {
	password := "samePassword"

	hash1, err := HashPassword(password)
	require.NoError(t, err)

	hash2, err := HashPassword(password)
	require.NoError(t, err)

	require.NotEqual(t, hash1, hash2)
}
