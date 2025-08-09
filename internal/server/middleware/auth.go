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
	cfg         config.Config
}

func (auth *userAuth) validate() error {
	var err error
	auth.token, err = jwt.ParseWithClaims(auth.tokenString, auth, func(t *jwt.Token) (interface{}, error) {
		return []byte(auth.cfg.SecretKey), nil
	})
	if err != nil {
		return err
	}
	if !auth.token.Valid {
		return errors.New("invalid token")
	}
	return nil
}
