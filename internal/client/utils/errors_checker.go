package utils

import (
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

// IsExpectedError check error is expected or not
func IsExpectedError(err error, expectedErrors []codes.Code) bool {
	statusErr, ok := status.FromError(err)
	if !ok {
		return false
	}
	for _, code := range expectedErrors {
		if statusErr.Code() == code {
			return true
		}
	}
	return false
}
