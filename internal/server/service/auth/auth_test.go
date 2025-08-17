package auth

import (
	"testing"
	"time"

	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

func TestCreateToken(t *testing.T) {
	cfg := &config.Config{
		SecretKey: "testsecret",
		TokenExp:  1 * time.Hour,
	}
	srv := NewService(cfg)
	login := "testuser"
	tokenStr, err := srv.CreateToken(login)
	require.NoError(t, err)
	require.NotEmpty(t, tokenStr)
	parsedToken, err := jwt.Parse(tokenStr, func(token *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})
	require.NoError(t, err)
	require.True(t, parsedToken.Valid)
	claims := parsedToken.Claims.(jwt.MapClaims)
	require.Equal(t, login, claims["Login"])
}
