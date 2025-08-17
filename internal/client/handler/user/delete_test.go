package user

import (
	"testing"

	mockUser "github.com/VladimirSh98/GophKeeper/mocks/user_service"
	"github.com/golang/mock/gomock"
	"github.com/spf13/cobra"
	"github.com/stretchr/testify/require"
)

func TestDeleteCmdRun(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	mockService := mockUser.NewMockServiceInterface(ctrl)
	handler := &Handler{
		userService: mockService,
	}

	cmd := handler.DeleteCmd()

	mockService.EXPECT().Delete(gomock.Any(), gomock.Any()).DoAndReturn(
		func(cmd *cobra.Command, args []string) {
			require.IsType(t, &cobra.Command{}, cmd)
		},
	)

	err := cmd.Execute()
	require.NoError(t, err)
}
