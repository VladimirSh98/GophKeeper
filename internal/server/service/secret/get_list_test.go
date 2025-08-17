package secret

import (
	"context"
	"encoding/json"
	"errors"
	"testing"
	"time"

	"github.com/VladimirSh98/GophKeeper/internal/server/repository/secret"
	"github.com/VladimirSh98/GophKeeper/internal/server/utils"
	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret"
	pb "github.com/VladimirSh98/GophKeeper/proto"
	"github.com/golang/mock/gomock"
	"github.com/stretchr/testify/require"
	"go.uber.org/zap"
)

func TestGetSecrets(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockSecretRepo := mockSecret.NewMockRepository(ctrl)
	logger, _ := zap.NewDevelopment()
	grpcSvc := NewSecretGrpc(nil, mockSecretRepo, logger)

	ctx := context.WithValue(context.Background(), utils.UserLoginKey, "user1")

	t.Run("successful get secrets", func(t *testing.T) {
		metadataBytes, _ := json.Marshal(map[string]string{"key": "value"})
		mockSecretRepo.EXPECT().
			GetSecretsByUser(ctx, "user1").
			Return([]secret.Secret{
				{
					ID:        1,
					UserID:    1,
					Archived:  false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DataType:  1,
					Content:   []byte("data"),
					Metadata:  metadataBytes,
				},
			}, nil)

		resp, err := grpcSvc.Get(ctx, &pb.GetSecretsRequest{})
		require.NoError(t, err)
		require.Len(t, resp.Secrets, 1)
		require.Equal(t, int64(1), resp.Secrets[0].Id)
		require.Equal(t, "value", resp.Secrets[0].Metadata["key"])
	})

	t.Run("repo error", func(t *testing.T) {
		mockSecretRepo.EXPECT().
			GetSecretsByUser(ctx, "user1").
			Return(nil, errors.New("db error"))

		resp, err := grpcSvc.Get(ctx, &pb.GetSecretsRequest{})
		require.Nil(t, resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "secret lookup failed")
	})

	t.Run("metadata unmarshal error", func(t *testing.T) {
		mockSecretRepo.EXPECT().
			GetSecretsByUser(ctx, "user1").
			Return([]secret.Secret{
				{
					ID:        1,
					UserID:    1,
					Archived:  false,
					CreatedAt: time.Now(),
					UpdatedAt: time.Now(),
					DataType:  1,
					Content:   []byte("data"),
					Metadata:  []byte{0xff, 0xfe}, // invalid JSON
				},
			}, nil)

		resp, err := grpcSvc.Get(ctx, &pb.GetSecretsRequest{})
		require.Nil(t, resp)
		require.Error(t, err)
		require.Contains(t, err.Error(), "secret metadata unmarshal failed")
	})
}
