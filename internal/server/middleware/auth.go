package middleware

import (
	"errors"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/golang-jwt/jwt/v4"
)

type userAuth struct {
	jwt.RegisteredClaims
	tokenString string
	token       *jwt.Token
	Login       string
}

func (auth *userAuth) validate(cfg *config.Config) error {
	var err error
	auth.token, err = jwt.ParseWithClaims(auth.tokenString, auth, func(t *jwt.Token) (interface{}, error) {
		return []byte(cfg.SecretKey), nil
	})
	if err != nil {
		return err
	}
	if !auth.token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
