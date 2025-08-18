package utils

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func TestIsExpectedError(t *testing.T) {
	t.Run("nil error", func(t *testing.T) {
		require.False(t, IsExpectedError(nil, []codes.Code{codes.Internal}))
	})

	t.Run("non-status error", func(t *testing.T) {
		err := errors.New("some error")
		require.False(t, IsExpectedError(err, []codes.Code{codes.Internal}))
	})

	t.Run("matching code", func(t *testing.T) {
		err := status.Error(codes.NotFound, "not found")
		require.True(t, IsExpectedError(err, []codes.Code{codes.NotFound, codes.Internal}))
	})

	t.Run("non-matching code", func(t *testing.T) {
		err := status.Error(codes.PermissionDenied, "denied")
		require.False(t, IsExpectedError(err, []codes.Code{codes.Internal, codes.NotFound}))
	})

	t.Run("empty expectedErrors", func(t *testing.T) {
		err := status.Error(codes.Internal, "internal")
		require.False(t, IsExpectedError(err, []codes.Code{}))
	})
}
