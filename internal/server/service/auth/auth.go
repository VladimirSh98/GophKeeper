package auth

import (
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/golang-jwt/jwt/v4"
	"time"
)

type auth struct {
	jwt.RegisteredClaims
	Login string
	cfg   *config.Config
}

// Service auth interface
type Service interface {
	CreateToken(login string) (string, error)
}

func NewService(cfg *config.Config) Service {
	return &auth{cfg: cfg}
}

// CreateToken create new token by login
func (s *auth) CreateToken(login string) (string, error) {
	exp := time.Hour * s.cfg.TokenExp
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, auth{
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(exp)),
		},
		Login: login,
	})
	tokenString, err := token.SignedString([]byte(s.cfg.SecretKey))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}
