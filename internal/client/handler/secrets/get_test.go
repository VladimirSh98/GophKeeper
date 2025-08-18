package secrets

import (
	"testing"

	mockSecret "github.com/VladimirSh98/GophKeeper/mocks/secret_service"
	"github.com/golang/mock/gomock"
)

func TestGetCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockSecret.NewMockServiceInterface(ctrl)
	handler := &Handler{
		secretService: mockService,
	}

	cmd := handler.GetCmd()

	mockService.EXPECT().Get(gomock.Any(), gomock.Any())

	err := cmd.Execute()
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
}
