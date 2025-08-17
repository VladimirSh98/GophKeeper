package middleware

import (
	"testing"
	"time"

	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
)

func TestUserAuth(t *testing.T) {
	cfg := &config.Config{
		SecretKey: "test-secret",
	}

	claims := &userAuth{
		Login: "testuser",
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(cfg.SecretKey))
	require.NoError(t, err)

	auth := &userAuth{tokenString: tokenString}
	err = auth.validate(cfg)
	require.NoError(t, err)
	require.Equal(t, "testuser", auth.Login)

	authInvalid := &userAuth{tokenString: "invalid.token.string"}
	err = authInvalid.validate(cfg)
	require.Error(t, err)
}
