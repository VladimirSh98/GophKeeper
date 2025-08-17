package middleware

import (
	"context"
	"testing"
	"time"

	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	"github.com/golang-jwt/jwt/v4"
	"github.com/stretchr/testify/require"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func TestJWTUnaryInterceptor(t *testing.T) {
	cfg := &config.Config{SecretKey: "test-secret"}

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

	interceptor := JWTUnaryInterceptor(cfg)

	ctx := metadata.NewIncomingContext(context.Background(), metadata.Pairs("authorization", tokenString))

	handler := func(ctx context.Context, req interface{}) (interface{}, error) {
		login := ctx.Value(utils.UserLoginKey)
		return login, nil
	}

	resp, err := interceptor(ctx, nil, &grpc.UnaryServerInfo{FullMethod: "/proto.User/GetSecret"}, handler)
	require.NoError(t, err)
	require.Equal(t, "testuser", resp)

	ctxNoToken := context.Background()
	_, err = interceptor(ctxNoToken, nil, &grpc.UnaryServerInfo{FullMethod: "/proto.User/GetSecret"}, handler)
	st, _ := status.FromError(err)
	require.Equal(t, codes.Unauthenticated, st.Code())

	ctxExcluded := context.Background()
	respExcluded, err := interceptor(ctxExcluded, nil, &grpc.UnaryServerInfo{FullMethod: "/proto.User/Login"}, handler)
	require.NoError(t, err)
	require.Nil(t, respExcluded)
}
