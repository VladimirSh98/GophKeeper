package middleware

import (
	"context"
	"github.com/VladimirSh98/GophKeeper/internal/server/config"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func JWTUnaryInterceptor(cfg *config.Config) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		if isExcludedMethod(info.FullMethod) {
			return handler(ctx, req)
		}

		md, ok := metadata.FromIncomingContext(ctx)
		if !ok {
			return nil, status.Error(codes.Unauthenticated, "metadata is not provided")
		}

		token := md.Get("authorization")
		if len(token) == 0 {
			return nil, status.Error(codes.Unauthenticated, "authorization token is not provided")
		}
		user := userAuth{tokenString: token[0]}
		err := user.validate(cfg)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "invalid token: %v", err)
		}

		newCtx := context.WithValue(ctx, utils.UserLoginKey, user.Login)
		return handler(newCtx, req)
	}
}

func isExcludedMethod(method string) bool {
	excluded := map[string]struct{}{
		"/proto.User/Login":    {},
		"/proto.User/Register": {},
	}
	_, ok := excluded[method]
	return ok
}
